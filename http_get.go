package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	name := html.EscapeString("<h1>Josemar</h1>")
	fmt.Println(name)
	resp, _ := http.Get("http://www.kindelbit.com")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
