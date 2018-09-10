package main

import (
	"github.com/scaat/kangaroo/container"
	"github.com/caicloud/cyclone/pkg/log"
	"os"
)

func Run(tty bool, command string) {
	parent := container.NewParentProcess(tty, command)
	if err := parent.Start(); err != nil {
		log.Error(err)
	}
	parent.Wait()
	os.Exit(-1)
// 	implement
}
