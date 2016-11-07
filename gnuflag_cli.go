package main

import (
	"fmt"

	flag "launchpad.net/gnuflag"
)

var name = flag.String("name", "World", "A name to say hello to")
var spanish bool

func init() {
	flag.BoolVar(&spanish, "spanish", false, "Use Spanish language")
	flag.BoolVar(&spanish, "s", false, "Use Spanish language")
}

func main() {
	flag.Parse(true)

	if spanish {
		fmt.Printf("Hola, %s!\n", *name)
	} else {
		fmt.Printf("Hello, %s!\n", *name)
	}
}
