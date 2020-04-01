package main

import "testing"

// https://talks.golang.org/2012/10things.slide#2
var groupedConfig struct {
	APIKey      string
	OAuthConfig []byte
}

func TestHelloWorld(t *testing.T) {
	groupedConfig.APIKey = "BADC0C0A"
	t.Log(groupedConfig.APIKey)
}
