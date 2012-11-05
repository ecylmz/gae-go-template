// Copyright 2012 Foo Bar <example@example.com>.

package app

import (
	"appengine"
	"appengine/user"
	"net/http"
	"text/template"
	"library/render"
	"bytes"
	// "fmt"
	// "errors"
)

func init() {
	http.HandleFunc("/", root)
}

func root(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	current_user := user.Current(c)

	passedTemplate := new(bytes.Buffer)
	template.Must(template.ParseFiles("templates/index.html")).Execute(passedTemplate, current_user)
	render.Render(w, r, passedTemplate)
}
