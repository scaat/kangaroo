package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strconv"
	"syscall"
	"github.com/urfave/cli"
	log "github.com/sirupsen/logrus"
)

const usage = `kangaroo is a simple container runtime implementation.
			   The purpose of this project is to learn how docker works and how to
write a docker by ourselves
			   Enjoy it, just for fun.`

func main() {

	app := cli.NewApp()
	app.Name = "kangaroo"
	app.Usage = usage

	app.Commands = []cli.Command{
		initCommand,
		runCommand,
	}

	app.Before = func(context *cli.Context) error {
		// Log as JSON instead of the default ASCII formatter.
		log.SetFormatter(&log.JSONFormatter{})

		log.SetOutput(os.Stdout)
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
