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
    Name string         `json:"name"`
    Value int           `json:"value"`
    ItemStyle ItemStyle `json:"itemStyle"`
}

type ItemStyle struct {
    Color string        `json:"color"`
}


func indexHandler(w http.ResponseWriter, r *http.Request) {
    data := []Data{
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
    }

    options := Options{
        Title: "Test",
        Data: data,
    }

    buf := &bytes.Buffer{}
    err := tpl.Execute(buf, options)
    if err != nil {
        log.Fatal(err)
    }

    buf.WriteTo(w)
}
