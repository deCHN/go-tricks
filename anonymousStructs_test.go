package main_test

import "testing"

type User struct{}

func TestTemplateData(t *testing.T) {
	data := struct {
		Title string
		Users []*User
	}{
		"chef",
		[]*User{&User{}, &User{}},
	}

	_ = data
}
