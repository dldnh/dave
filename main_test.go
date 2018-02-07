package main

import "testing"

func test1(t *testing.T) {
	name := getName()
	expect := "Dave"
	if name != expect {
		t.Fatalf("got incorrect name: %s", name, expect)
	}
}
