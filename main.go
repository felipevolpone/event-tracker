package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "event tracker across your AWS infraestructure"
	app.Usage = "cli"

	app.Commands = []cli.Command{
		snsTrack(),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func snsTrack() cli.Command {
	return cli.Command{
		Name:   "sns",
		Usage:  "start the track from a given sns arn",
		Action: SNSTrack,
	}
}
