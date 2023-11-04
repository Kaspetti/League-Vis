package server

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
)


var tpl = template.Must(template.ParseFiles("pages/index.html"))


type Options struct {
    Title string
    Data []Data
}


type Data struct {
    Name string
    Value int
    ItemStyle ItemStyle
}

type ItemStyle struct {
    Color string
}


func indexHandler(w http.ResponseWriter, r *http.Request) {
    options := Options{
        Title: "Test",
        Data: []Data{
            {
                Name: "Test Champ 0",
                Value: 15,
                ItemStyle: ItemStyle{
                    Color: "#FF0000",
                },
            },
            {
                Name: "Test Champ 1",
                Value: 15,
                ItemStyle: ItemStyle{
                    Color: "#00FF00",
                },
            },
            {
                Name: "Test Champ 2",
                Value: 15,
                ItemStyle: ItemStyle{
                    Color: "#0000FF",
                },
            },
        },
    }

    buf := &bytes.Buffer{}
    err := tpl.Execute(buf, options)
    if err != nil {
        log.Fatal(err)
    }

    buf.WriteTo(w)
}
