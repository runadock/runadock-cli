package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

// NewKillCommand creates a new KilCommand with all Arguments.
func NewKillCommand() cli.Command {
	return cli.Command{
		Name:  "kill",
		Usage: "kill a container",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "container, id", Usage: "ID of the container to kill"},
		},
		Action: func(c *cli.Context) {
			killCommandFunc(c)
		},
	}
}

func killCommandFunc(c *cli.Context) {
	id := c.String("container")
	if id == "" {
		fmt.Println("container may not be empty.")
		os.Exit(1)
	}

	url := baseurl + "/api/v1/container/" + id

	body := delete(url)
	fmt.Println(string(body[:]))
}
