package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

// NewRunCommand creates a new RunCommand with all Arguments.
func NewRunCommand() cli.Command {
	return cli.Command{
		Name:  "run",
		Usage: "run a new container",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "source, src", Usage: "the source to the Container"},
			cli.StringFlag{Name: "name", Usage: "the name for the container"},
			cli.StringFlag{Name: "size", Usage: "the size of the container"},
			cli.StringFlag{Name: "plan", Usage: "the plan for the container"},
		},
		Action: func(c *cli.Context) {
			runCommandFunc(c)
		},
	}
}

func runCommandFunc(c *cli.Context) {
	if c.String("source") == "" {
		fmt.Println("source my not be empty.")
		os.Exit(1)
	}
	runContainer := &createContainerRequest{
		Source: c.String("source"),
		Name:   c.String("name"),
		Size:   c.String("size"),
		Plan:   c.String("plan"),
	}
	runContainerJSON, _ := json.Marshal(runContainer)
	url := baseurl + "/api/v1/container"

	body := post(url, runContainerJSON)
	var container Container
	err := json.Unmarshal(body, &container)
	if err != nil {
		panic("Unable to deserialize describe container response.")
	}
	fmt.Println("created:", container.ID)
}
