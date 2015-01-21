package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/codegangsta/cli"
	"github.com/ryanuber/columnize"
)

// NewPSCommand creates a new PSCommand with all Arguments.
func NewPSCommand() cli.Command {
	return cli.Command{
		Name:  "ps",
		Usage: "show all container",
		Flags: []cli.Flag{
			cli.BoolFlag{Name: "detail, d", Usage: "show long container ids."},
			cli.BoolFlag{Name: "all, a", Usage: "show even terminated container's."},
		},
		Action: func(c *cli.Context) {
			psCommandFunc(c)
		},
	}
}

func psCommandFunc(c *cli.Context) {
	url := baseurl + "/api/v1/container"

	detail := c.Bool("detail")
	all := c.Bool("all")
	if all {
		url = url +"?all=true"
	}

	body := get(url)

	var containers []Container
	err := json.Unmarshal(body, &containers)
	if err != nil {
		fmt.Println("Respone body:" + string(body))
		panic("Unable to deserialize describe container response.")
	}

	output := []string{
		"ContainerID | SOURCE | CREATED | STATUS | DNS | PORTS | NAME",
	}

	for i := range containers {
		container := containers[i]
		ports := ""
		if len(container.Ports) > 0 {
			port := container.Ports[0]
			ports = fmt.Sprintf("%d->%d/%s", port.PublicPort, port.PrivatePort, port.Scheme)
		}
		containerID := ""
		if len(container.ContainerID) > 0 {
			if detail {
				containerID = container.ContainerID
			} else {
				containerID = container.ContainerID[:11]
			}
		}
		created := time.Unix(container.Created/1000, 0)
		output = append(output, containerID+" | "+
			container.Source+" | "+
			string(created.Format("2006-01-02 15:04:05"))+" | "+
			container.State+" | "+
			container.PublicDNS+" | "+
			ports+" | "+
			container.Name)
	}
	result := columnize.SimpleFormat(output)
	fmt.Println(result)

}
