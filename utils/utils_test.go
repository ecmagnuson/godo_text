package utils

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

//test non exported functions in utils.go

func setup() (string, int) {
	return "(8675309) something (8) @home", 8675309
}

func Test_getId(t *testing.T) {
	s, want := setup()
	if getID(s) != want {
		t.Fatalf("getId failed to parse correct ID")
	}
}

func Test_hasContext(t *testing.T) {
	s, _ := setup()
	if !hasContext(s) {
		t.Fatalf("hasContext failed to find a context when there was one")
	}
}
