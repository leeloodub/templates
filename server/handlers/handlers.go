package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func PutEvent(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	fo, err := os.OpenFile("some-data-file", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	if _, err := fo.Write(b); err != nil {
		panic(err)
	}
}

func Acknowledge(w http.ResponseWriter, r *http.Request) {
	m := r.URL.Path
	m = strings.TrimPrefix(m, "/")
	m = "Hello " + m
	w.Write([]byte(m))
}

func ShowIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Print("index page")
}
