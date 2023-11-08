package server

import (
	"html/template"
	"net/http"
    "bytes"
    "log"
)


var indexTpl = template.Must(template.ParseFiles("pages/index.html"))


func indexHandler(w http.ResponseWriter, r *http.Request) {
    type Options struct{}

    buf := &bytes.Buffer{}
    err := indexTpl.Execute(buf, Options{})
    if err != nil {
        log.Fatal(err)
    }

    buf.WriteTo(w)
}

