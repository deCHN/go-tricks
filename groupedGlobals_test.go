package main

import "testing"

var groupedConfig struct {
	APIKey      string
	OAuthConfig []byte
}

func TestHelloWorld(t *testing.T) {
	groupedConfig.APIKey = "BADC0C0A"
	t.Log(groupedConfig.APIKey)
}
