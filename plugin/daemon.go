package plugin

import (
	"io/ioutil"
	"os"
)

const (
	dockerExe  = "/usr/local/bin/docker"
	dockerdExe = "/usr/local/bin/dockerd"
	dockerHome = "/root/.docker/"
)

func (p Plugin) startDaemon() {
	cmd := commandDaemon(p.settings.Daemon)
	if p.settings.Daemon.Debug {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	} else {
		cmd.Stdout = ioutil.Discard
		cmd.Stderr = ioutil.Discard
	}
	go func() {
		trace(cmd)
		_ = cmd.Run()
	}()
}
