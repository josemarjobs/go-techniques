// application to count up and down, a simple showcase for the package cli
package main

import (
	"fmt"
	"os"

	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Usage = "Count up or down"
	app.Commands = []cli.Command{
		{
			Name:      "up",
			ShortName: "u",
			Usage:     "Count up",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "stop, s",
					Usage: "Value to count up to.",
					Value: 10,
				},
			},
			Action: func(c *cli.Context) error {
				stop := c.Int("stop")
				if stop <= 0 {
					fmt.Println("Stop has to be greater than 0.")
				}
				for i := 1; i <= stop; i++ {
					fmt.Println(i)
				}
				return nil
			},
		},
		{
			Name:      "down",
			ShortName: "d",
			Usage:     "Count down",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "start, s",
					Usage: "Start countdown from",
					Value: 10,
				},
			},
			Action: func(c *cli.Context) error {
				start := c.Int("start")
				if start <= 0 {
					fmt.Println("Start has to be greater than 0")
				}
				for i := start; i >= 1; i-- {
					fmt.Println(i)
				}
				return nil
			},
		},
	}

	app.Run(os.Args)
}
