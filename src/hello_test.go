package main

import "testing"

func TestHello(t *testing.T) {
    got := Hello("SciFiddy")
    want := "Hello, SciFiddy"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}