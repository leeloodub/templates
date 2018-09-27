package handlers

import (
	"os"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	// t.Fatal("not implemented")
}

type testCase struct {
	description string
	in          int
	expected    int
	err         error
}

var testCases = []testCase{
	testCase{
		description: "",
		in:          1,
		expected:    2,
		err:         nil,
	},
}

func TestBasic(t *testing.T) {
	for _, tt := range testCases {
		actual := replaceFunc(tt.in)
		if actual != tt.expected {
			t.Errorf("Test case %s replaceFunc(%v): expected %v, actual %v", tt.description, tt.in, tt.expected, actual)
		}
	}
}

func TestMain(m *testing.M) {
	replaceSetup()
	code := m.Run()
	replaceTeardown()
	os.Exit(code)
}
