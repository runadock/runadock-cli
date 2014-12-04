package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

// NewInspectCommand creates a new InspectCommand with all Arguments.
func NewInspectCommand() cli.Command {
	return cli.Command{
		Name:  "inspect",
		Usage: "inspect a container",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "container, id", Usage: "ID of the container to inspect"},
		},
		Action: func(c *cli.Context) {
			inspectCommandFunc(c)
		},
	}
}

func inspectCommandFunc(c *cli.Context) {
	id := c.String("container")
	if id == "" {
		fmt.Println("container may not be empty.")
		os.Exit(1)
	}

	url := baseurl + "/api/v1/container/" + id

	body := get(url)
	var container Container
	err := json.Unmarshal(body, &container)
	if err != nil {
		panic("Unable to deserialize describe container response.")
	}
	j, jerr := json.MarshalIndent(container, "", "  ")
	if jerr != nil {
		fmt.Println("jerr:", jerr.Error())
	}
	fmt.Println(string(j))
}
