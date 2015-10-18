package main

import (
	"bytes"
	"html/template"
	"testing"
)

type User struct {
	FirstName      string
	FavoriteColors []string
}

func BenchmarkDefaultTemplate(b *testing.B) {
	var buf bytes.Buffer
	t, err := template.New("index.html").ParseFiles("index.html")
	myUser := &User{
		FirstName:      "Bob",
		FavoriteColors: []string{"blue", "green", "mauve"},
	}
	for i := 0; i < b.N; i++ {
		err := template.Must(t, err).Execute(&buf, myUser)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkEgoTemplate(b *testing.B) {
	var buf bytes.Buffer
	myUser := &User{
		FirstName:      "Bob",
		FavoriteColors: []string{"blue", "green", "mauve"},
	}
	for i := 0; i < b.N; i++ {
		err := MyTmpl(&buf, myUser)
		if err != nil {
			panic(err)
		}
	}
}
