package main

import (
	"bootinspect/inspector"
	"fmt"
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
		fmt.Println("Hello friend!")
		fmt.Println(c.Args().First())
		inspector.DdoThis()
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
