package main

import (
	"fmt"
	mycli "github.com/jw3/ppc/cli"
	"github.com/urfave/cli"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	app := cli.NewApp()
	cfg := mycli.NewConfiguration()

	app.Commands = []*cli.Command{
		{
			Name:        "call",
			Usage:       "Call a particular function on a device",
			UsageText:   "polyform call [options] <device> <function> [argument]",
			Description: "Call a particular function on a device",
			ArgsUsage:   "[options] <device> <function> [argument]",
			Action: func(c *cli.Context) error {
				d := c.Args().Get(0)
				f := c.Args().Get(1)
				uri := fmt.Sprintf("http://%s/devices/%s/%s", cfg.ApiUri, d, f)
				println(uri)

				// POST /v1/devices/{DEVICE_ID}/{FUNCTION}
				a := c.Args().Get(2)
				v := url.Values{}
				if len(a) > 0 {
					v.Set("args", a)
				}
				_, e := http.PostForm(uri, v)

				return e
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
