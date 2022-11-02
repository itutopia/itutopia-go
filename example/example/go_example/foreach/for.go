package foreach

import (
	"os"
	"testing"
)

func TestFor(t *testing.T) {
	for true {
		t.Log("forLoop")
		os.Exit(1)
	}
}

func TestWhileLoop(t *testing.T) {
	n := 0
	// while(n<5)
	for n < 5 {
		t.Log(n)
		n++
	}
}
