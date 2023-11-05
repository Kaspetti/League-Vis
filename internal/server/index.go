package server

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"math"
	"net/http"
	"os"

	"github.com/Kaspetti/League-Vis/internal/datahandling"
)


var tpl = template.Must(template.ParseFiles("pages/index.html"))


type Options struct {
    Title string
    TotalPlayed float64
    Data []Data 
}


type Data struct {
    Name string         `json:"name"`
    Value float64           `json:"value"`
    ItemStyle ItemStyle `json:"itemStyle"`
}

type ItemStyle struct {
    Color string        `json:"color"`
}


func indexHandler(w http.ResponseWriter, r *http.Request) {
    champ := os.Args[1]

    totalPlayed := 0.
    champStats := datahandling.GetChampionStats(champ, BotlaneData)
    data := make([]Data, 0)
    for name, stats := range champStats {
        d := Data{
            Name: name,
            Value: stats.Played,
            ItemStyle: ItemStyle{
                Color: interpolateColor(stats.Winrate, 45, 50, 55, 0.8),
            },
        }

        totalPlayed += stats.Played
        data = append(data, d)
    }

    options := Options{
        Title: fmt.Sprintf("%s Synergies", champ),
        TotalPlayed: totalPlayed,
        Data: data,
    }

    buf := &bytes.Buffer{}
    err := tpl.Execute(buf, options)
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
		green = greenColor
		blue = (1 - normalized) * blueColor
	}

	// Convert to hex color
	hex := fmt.Sprintf("#%02X%02X%02X", int(math.Round(red * brightness)), int(math.Round(green * brightness)), int(math.Round(blue * brightness)))
	return hex
}
