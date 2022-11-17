package utils

import (
	"log"
	"testing"
)

//test non exported functions in utils.go

var cases = []struct {
	s       string
	want    int
	context bool
}{
	{"(8675309) something (8) @home", 8675309, true},
	{"(1) something (8) @home", 1, true},
	{"(8675309) something (8)", 8675309, false},
}

func Test_getId(t *testing.T) {
	for _, c := range cases {
		got := getID(c.s)
		if got != c.want {
			log.Fatalf("getID failed.")
		}
	}
}

func Test_hasContext(t *testing.T) {
	for _, c := range cases {
		got := hasContext(c.s)
		if got != c.context {
			log.Fatalf("hasContext failed.")
		}
	}
}
