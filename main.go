package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "mono-proto"
	app.Usage = "orchestrate protobuf generation from a monorepo"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Hello, world!")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
