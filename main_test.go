package main

import "testing"

func Test1(t *testing.T) {
	name := getName()
	expect := "Dave"
	if name != expect {
		t.Fatalf("got incorrect name: '%s' (expected: '%s')", name, expect)
	}
}
