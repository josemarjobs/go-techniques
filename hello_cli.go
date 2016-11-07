package main

import (
	"fmt"
	"os"

	cli "gopkg.in/urfave/cli.v1"
)

func application(c *cli.Context) error {
	name := c.GlobalString("name")
	fmt.Printf("Hello, %s\n", name)
	return nil
}
func main() {
	app := cli.NewApp()
	app.Name = "Hello"
	app.Usage = "Print Hello World"
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "name", Value: "World", Usage: "Who to say hello to."},
	}

	app.Action = func(c *cli.Context) error {

		name := c.GlobalString("name")
		fmt.Printf("Hello, %s\n", name)
		return nil
	}

	app.Run(os.Args)
}
