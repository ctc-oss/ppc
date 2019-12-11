package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	app := cli.NewApp()

	app.Commands = []*cli.Command{
		{
			Name:        "call",
			Usage:       "Call a particular function on a device",
			UsageText:   "particle call [options] <device> <function> [argument]",
			Description: "Call a particular function on a device",
			ArgsUsage:   "[options] <device> <function> [argument]",
			Action:  func(c *cli.Context) error {
				d := c.Args().Get(0)
				f := c.Args().Get(1)
				uri := fmt.Sprintf("http://localhost:9000/v1/devices/%s/%s", d, f)
				println(uri)

				// POST /v1/devices/{DEVICE_ID}/{FUNCTION}
				_, e := http.Post(uri, "text/plain", strings.NewReader(""))
				return e
			},
		},
	}



	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
