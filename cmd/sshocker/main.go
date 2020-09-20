package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func newApp() *cli.App {
	debug := false
	app := cli.NewApp()
	app.Name = "sshocker"
	app.Usage = "ssh + reverse sshfs + port forwarder, in Docker-like CLI"
	app.UsageText = "sshocker run -p LOCALPORT:REMOTEPORT -v LOCALDIR:REMOTEDIR USER@HOST"

	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:        "debug",
			Usage:       "debug mode",
			Destination: &debug,
		},
	}
	app.Flags = append(app.Flags, setHidden(runFlags, true)...)
	app.Before = func(context *cli.Context) error {
		if debug {
			logrus.SetLevel(logrus.DebugLevel)
		}
		return nil
	}
	app.Commands = []*cli.Command{runCommand}
	app.Action = runAction
	return app
}

func main() {
	if err := newApp().Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
