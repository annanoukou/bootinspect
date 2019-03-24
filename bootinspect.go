package main

import (
	"bootinspect/inspector"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Bootinspect"
	app.Usage = "Inspect CSS files"
	app.Author = "Blackpeak"
	app.Version = "0.0.1"

	app.Action = func(c *cli.Context) error {
		inspector.InspectFile(c.Args().First())
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
