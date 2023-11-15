package server

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"math"
	"net/http"
	"strings"

	"github.com/Kaspetti/League-Vis/internal/datahandling"
)


type Options struct {
    Title string
    TotalPlayed float64
    Green string
    White string
    Red string
    Data []Data 
}

const (
    SATURATION = 0.8
)

type Data struct {
    Name string         `json:"name"`
    Value []float64           `json:"value"`
    ItemStyle ItemStyle `json:"itemStyle"`
}

type ItemStyle struct {
    Color string        `json:"color"`
}


var championPageTpl = template.Must(template.ParseFiles("pages/championPage.html"))


func championPageHandler(w http.ResponseWriter, r *http.Request) {
    pathSegments := strings.Split(r.URL.Path, "/")
    if len(pathSegments) < 3 || pathSegments[2] == "" {
        http.NotFound(w, r)
        return
    }
    champion := pathSegments[2]

    totalPlayed := 0.
    champStats := datahandling.GetChampionStats(champion, BotlaneData)

    if len(champStats) == 0 {
        http.NotFound(w, r)
        return
    }

    data := make([]Data, 0)
    for name, stats := range champStats {
        d := Data{
            Name: strings.Title(name),
            Value: []float64{stats.Played, stats.Winrate},
            ItemStyle: ItemStyle{
                Color: interpolateColor(stats.Winrate, 40, 50, 60, SATURATION),
            },
        }

        totalPlayed += stats.Played
        data = append(data, d)
    }

    options := Options{
        Title: fmt.Sprintf("%s Synergies", strings.Title(champion)),
        TotalPlayed: totalPlayed,
        Green: interpolateColor(60, 40, 50, 60, SATURATION),
        White: interpolateColor(50, 40, 50, 60, SATURATION),
        Red: interpolateColor(40, 40, 50, 60, SATURATION),
        Data: data,
    }

    buf := &bytes.Buffer{}
    err := championPageTpl.Execute(buf, options)
    if err != nil {
        log.Fatal(err)
    }

    buf.WriteTo(w)
}


func interpolateColor(value, min, mid, max, brightness float64) string {
	// Ensure value is within the range
	if value < min {
		value = min
	} else if value > max {
		value = max
	}

	// Define the color components for red, white/grey, and green
	redColor := 255.0
	greenColor := 255.0
	blueColor := 255.0 // White if 255, grey if less (e.g., 220 for light grey)

	var red, green, blue float64

	if value < mid {
		// Interpolate from red to white/grey
		normalized := (value - min) / (mid - min)
		red = redColor
		green = normalized * greenColor
		blue = normalized * blueColor
	} else {
		// Interpolate from white/grey to green
		normalized := (value - mid) / (max - mid)
		red = (1 - normalized) * redColor
		green = (1 - normalized) * greenColor
		blue = blueColor
	}

	// Convert to hex color
	hex := fmt.Sprintf("#%02X%02X%02X", int(math.Round(red * brightness)), int(math.Round(green * brightness)), int(math.Round(blue * brightness)))
	return hex
}
