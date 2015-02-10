package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	"github.com/codegangsta/cli"
)

var (
	baseurl         = "https://runadock.io"
	user            string
	apiToken        string
	revision        string
	transportConfig *http.Transport
	insecure        = false
)

func main() {
	app := cli.NewApp()
	app.Name = "runadock "
	app.Version = "git sha1: " + revision
	app.Usage = "manage your containers at runadock.io"
	app.Flags = []cli.Flag{
		cli.BoolFlag{Name: "insecure", Usage: "allow connection to insecure ssl servers, not recommended."},
	}

	app.Commands = []cli.Command{
		NewRunCommand(),
		NewPSCommand(),
		NewKillCommand(),
		NewInspectCommand(),
	}

	user = os.Getenv("RUNADOCK_USER")
	apiToken = os.Getenv("RUNADOCK_TOKEN")
	newBaseUrl := os.Getenv("RUNADOCK_BASEURL")

	if newBaseUrl != "" {
		baseurl = newBaseUrl
	}

	if user == "" {
		fmt.Println("No user given. You must set the Environment Variable RUNADOCK_USER")
		fmt.Println("If you dont hava a runadock.io user already, signup at https://runadock.io .")
		os.Exit(1)
	}
	if apiToken == "" {
		fmt.Println("No API Token given. You must set the Environment Variable RUNADOCK_TOKEN.")
		fmt.Println("If you dont hava a token already, create one at https://runadock.io .")
		os.Exit(1)
	}

	// FIXME need to be configured from the insecure flag
	transportConfig = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	app.Run(os.Args)
}
