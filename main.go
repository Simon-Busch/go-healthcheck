package main

import (
	"github.com/urfave/cli/v2"
	"os"
	"log"
	"fmt"
)

func main() {
	fmt.Println("Starting")

	app := &cli.App {
		Name: "Health checker",
		Usage: "A tiny tool to check whether a website is running or is down",
		Flags: []cli.Flag {
			&cli.StringFlag {
				Name: "domain",
				Aliases: []string{"d"},
				Usage: "Domain name to check",
				Required: true,
			},
			&cli.StringFlag {
				Name: "port",
				Aliases: []string{"p"},
				Usage: "Port number to check",
				Required: false,
			},
		},
		Action: func(c *cli.Context) error {
			port := c.String("port")
			if c.String("port") == "" {
				port = "80"
			}

			status := Check(c.String("domain"), port)

			fmt.Println("Status: ", status)
			fmt.Println("Port: ", port)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
