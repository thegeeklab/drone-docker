package plugin

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

// Daemon defines Docker daemon parameters.
type Daemon struct {
	Registry      string          // Docker registry
	Mirror        string          // Docker registry mirror
	Insecure      bool            // Docker daemon enable insecure registries
	StorageDriver string          // Docker daemon storage driver
	StoragePath   string          // Docker daemon storage path
	Disabled      bool            // DOcker daemon is disabled (already running)
	Debug         bool            // Docker daemon started in debug mode
	Bip           string          // Docker daemon network bridge IP address
	DNS           cli.StringSlice // Docker daemon dns server
	DNSSearch     cli.StringSlice // Docker daemon dns search domain
	MTU           string          // Docker daemon mtu setting
	IPv6          bool            // Docker daemon IPv6 networking
	Experimental  bool            // Docker daemon enable experimental mode
}

// Login defines Docker login parameters.
type Login struct {
	Registry string // Docker registry address
	Username string // Docker registry username
	Password string // Docker registry password
	Email    string // Docker registry email
	Config   string // Docker Auth Config
}

// Build defines Docker build parameters.
type Build struct {
	Remote     string          // Git remote URL
	Name       string          // Git commit sha used as docker default named tag
	Ref        string          // Git commit ref
	Branch     string          // Git repository branch
	Dockerfile string          // Docker build Dockerfile
	Context    string          // Docker build context
	TagsAuto   bool            // Docker build auto tag
	TagsSuffix string          // Docker build tags with suffix
	Tags       cli.StringSlice // Docker build tags
	Args       cli.StringSlice // Docker build args
	ArgsEnv    cli.StringSlice // Docker build args from env
	Target     string          // Docker build target
	Pull       bool            // Docker build pull
	CacheFrom  cli.StringSlice // Docker build cache-from
	Compress   bool            // Docker build compress
	Repo       string          // Docker build repository
	NoCache    bool            // Docker build no-cache
	AddHost    cli.StringSlice // Docker build add-host
	Quiet      bool            // Docker build quiet
}

// Settings for the Plugin.
type Settings struct {
	Daemon  Daemon
	Login   Login
	Build   Build
	Dryrun  bool
	Cleanup bool
}

// Validate handles the settings validation of the plugin.
func (p *Plugin) Validate() error {
	p.settings.Build.Name = "00000000"
	if p.pipeline.Commit.SHA != "" {
		p.settings.Build.Name = p.pipeline.Commit.SHA
	}

	p.settings.Build.Branch = p.pipeline.Repo.Branch
	p.settings.Build.Ref = p.pipeline.Commit.Ref
	p.settings.Daemon.Registry = p.settings.Login.Registry

	if p.settings.Build.TagsAuto {
		// return true if tag event or default branch
		if UseDefaultTag(
			p.settings.Build.Ref,
			p.settings.Build.Branch,
		) {
			tag, err := DefaultTagSuffix(
				p.settings.Build.Ref,
				p.settings.Build.TagsSuffix,
			)
			if err != nil {
				logrus.Printf("cannot generate tags from %s, invalid semantic version", p.settings.Build.Ref)
				return err
			}
			p.settings.Build.Tags = *cli.NewStringSlice(tag...)
		} else {
			logrus.Printf("skip auto-tagging for %s, not on default branch or tag", p.settings.Build.Ref)
			return nil
		}
	}

	return nil
}

// Execute provides the implementation of the plugin.
func (p *Plugin) Execute() error {
	// start the Docker daemon server
	if !p.settings.Daemon.Disabled {
		p.startDaemon()
	}

	// poll the docker daemon until it is started. This ensures the daemon is
	// ready to accept connections before we proceed.
	for i := 0; i < 15; i++ {
		cmd := commandInfo()
		err := cmd.Run()
		if err == nil {
			break
		}
		time.Sleep(time.Second * 1)
	}

	// Create Auth Config File
	if p.settings.Login.Config != "" {
		if err := os.MkdirAll(dockerHome, 0o600); err != nil {
			return fmt.Errorf("failed to create docker home: %s", err)
		}

		path := filepath.Join(dockerHome, "config.json")
		err := ioutil.WriteFile(path, []byte(p.settings.Login.Config), 0o600)
		if err != nil {
			return fmt.Errorf("error writing config.json: %s", err)
		}
	}

	// login to the Docker registry
	if p.settings.Login.Password != "" {
		cmd := commandLogin(p.settings.Login)
		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("error authenticating: %s", err)
		}
	}

	switch {
	case p.settings.Login.Password != "":
		fmt.Println("Detected registry credentials")
	case p.settings.Login.Config != "":
		fmt.Println("Detected registry credentials file")
	default:
		fmt.Println("Registry credentials or Docker config not provided. Guest mode enabled.")
	}

	// add proxy build args
	addProxyBuildArgs(&p.settings.Build)

	var cmds []*exec.Cmd
	cmds = append(cmds, commandVersion()) // docker version
	cmds = append(cmds, commandInfo())    // docker info

	// pre-pull cache images
	for _, img := range p.settings.Build.CacheFrom.Value() {
		cmds = append(cmds, commandPull(img))
	}

	cmds = append(cmds, commandBuild(p.settings.Build)) // docker build

	for _, tag := range p.settings.Build.Tags.Value() {
		cmds = append(cmds, commandTag(p.settings.Build, tag)) // docker tag

		if !p.settings.Dryrun {
			cmds = append(cmds, commandPush(p.settings.Build, tag)) // docker push
		}
	}

	if p.settings.Cleanup {
		cmds = append(cmds, commandRmi(p.settings.Build.Name)) // docker rmi
		cmds = append(cmds, commandPrune())                    // docker system prune -f
	}

	// execute all commands in batch mode.
	for _, cmd := range cmds {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		trace(cmd)

		err := cmd.Run()

		switch {
		case err != nil && isCommandPull(cmd.Args):
			fmt.Printf("Could not pull cache-from image %s. Ignoring...\n", cmd.Args[2])
		case err != nil && isCommandPrune(cmd.Args):
			fmt.Printf("Could not prune system containers. Ignoring...\n")
		case err != nil && isCommandRmi(cmd.Args):
			fmt.Printf("Could not remove image %s. Ignoring...\n", cmd.Args[2])
		default:
			return err
		}
	}

	return nil
}
