package main

import (
	"os"

	"github.com/scaat/kangaroo/container"
	"github.com/sirupsen/logrus"
)

func Run(tty bool, command string) {
	parent := container.NewParentProcess(tty, command)
	if err := parent.Start(); err != nil {
		logrus.Error(err)
	}
	parent.Wait()
	os.Exit(-1)
}
