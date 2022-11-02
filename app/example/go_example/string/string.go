package string

import "testing"

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*" + s)
	t.Log(len(s))
	t.Logf(string(len(s)))
}
