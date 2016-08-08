package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: urlencode <path fragment>")
	}
	u := &url.URL{Path: strings.Join(os.Args[1:], " ")}
	fmt.Println(u.String())
}
