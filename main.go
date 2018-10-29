package main

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/blizztrack/ribbit-go"
	"github.com/sanity-io/litter"

	"github.com/urfave/cli"
)

var (
	region string
)

func process(c *cli.Context) error {
	var game string
	var mode string
	if len(c.Args()) > 1 {
		mode = strings.ToLower(c.Args()[1])
		game = strings.ToLower(c.Args()[0])
	} else {
		game = strings.ToLower(c.Args()[0])
	}

	client := ribbit.NewRibbitClient(region)
	if game == "summary" {
		summary, err := client.Summary()
		if err != nil {
			return err
		}

		litter.Dump(summary)
		return nil
	}

	var res interface{}
	switch mode {
	case "versions":
		vers, err := client.Versions(game)
		if err != nil {
			return err
		}
		res = vers
		break
	case "bgdl":
		vers, err := client.BGDL(game)
		if err != nil {
			return err
		}
		res = vers
		break
	default:
		return errors.New("supported modes: versions, bgdl")
	}

	litter.Dump(res)
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "Ribbit Cli Client"
	app.Description = "Cli Client for access Battle.net Ribbit services"
	app.HideVersion = true

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "region, r",
			Value:       "us",
			Usage:       "Region to use: us, eu, kr, sg, cn",
			Destination: &region,
		},
	}

	app.Action = process

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
