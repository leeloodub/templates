package main

import (
	"fmt"
	"log"
	"net/url"

	myservice "./myservice"
)

var (
	client *myservice.Client
)

func main() {
	u, err := url.Parse("http://domain.com/search?q=foo")
	if err != nil {
		log.Fatal(err)
	}
	u.Scheme = "https"
	client = myservice.NewClient(u)
	fmt.Print("hey")
}
