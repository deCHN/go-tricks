package main

import "testing"

type User struct {}
func TestTemplateData(t *testing.T) {
	data := struct {
		Title string
		Users []*{}User
	}{
		title,
		users,
	}
}
