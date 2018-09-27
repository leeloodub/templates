package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	url     = "http://replace"
	port    = ":8080"
	putPath = "replace"
)

type Entry struct {
	Field string `json:"field"`
}

func main() {
	content, err := ioutil.ReadFile("data-file")
	if err != nil {
		fmt.Printf("could not read content of file: %v", err)
	}

	entries := []Entry{}
	err = json.Unmarshal(content, &entries)
	if err != nil {
		fmt.Printf("could not parse file: %v", err)
	}

	for _, e := range entries {
		encoded, _ := json.Marshal(e)
		resp, err := http.Post(url+port+putPath, "application/json", bytes.NewBuffer(encoded))
		if err != nil {
			fmt.Print(err)
		}
		fmt.Print(resp)
	}
}
