package mascot_test

import (
	"testing"

	"github.com/nickalara/first-go-app/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Go Gopher" {
		t.Fatal("Wrong mascot...")
	}
}
