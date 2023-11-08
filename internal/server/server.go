package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Kaspetti/League-Vis/internal/datahandling"
)


var BotlaneData []datahandling.BotlaneData


func RunServer(ip, port string) {
    var err error
    BotlaneData, err = datahandling.ImportData("botlanes.csv")
    if err != nil {
        log.Fatalln(err)
    }


    fs := http.FileServer(http.Dir("assets"))

    mux := http.NewServeMux()

    mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
    mux.HandleFunc("/", indexHandler)
    mux.HandleFunc("/champions/", championPageHandler)

    http.ListenAndServe(fmt.Sprintf("%s:%s", ip, port), mux)
}
