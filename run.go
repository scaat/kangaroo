package main

import (
	"os"

	"strings"

	"github.com/scaat/kangaroo/container"
	"github.com/sirupsen/logrus"
)

func Run(tty bool, comArray []string) {
	parent, writePipe := container.NewParentProcess(tty)
	if parent == nil {
		logrus.Errorf("New parent process error")
		return
	}
	if err := parent.Start(); err != nil {
		logrus.Error(err)
	}
	// cgroupManager := cgroups.NewCgroupManager("kangaroo-cgroup")
	// defer cgroupManager.Destroy()

	// cgroupManager.Set(res)
	// cgroupManager.Apply(parent.Process.Pid)
	sendInitCommand(comArray, writePipe)
	parent.Wait()
}

func sendInitCommand(comArray []string, writePipe *os.File) {
	command := strings.Join(comArray, " ")
	logrus.Infof("command all is %s", command)
	writePipe.WriteString(command)
	writePipe.Close()
}
