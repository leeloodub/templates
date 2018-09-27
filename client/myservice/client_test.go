package myservice

import (
	"os"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	// t.Fatal("not implemented")
}

func TestMain(m *testing.M) {
	// replaceSetup()
	code := m.Run()
	// replaceTeardown()
	os.Exit(code)
}
