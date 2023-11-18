package server

import (
	"fmt"
	"math"
	"net/http"
	"strings"

	"github.com/Kaspetti/League-Vis/internal/datahandling"
	"github.com/gin-gonic/gin"
)


const (
    SATURATION = 0.8
)


type Options struct {
    Title string            `json:"title"`
    TotalPlayed float64     `json:"totalPlayed"`
    WinColor string         `json:"winColor"`
    NeutralColor string     `json:"neutralColor"`
    LoseColor string        `json:"loseColor"`
    Data []Data             `json:"data"` 
}


type Data struct {
    Name string         `json:"name"`
    Value []float64     `json:"value"`
    ItemStyle ItemStyle `json:"itemStyle"`
    Label Label         `json:"label"`
}


type ItemStyle struct {
    Color string        `json:"color"`
}


type Label struct {
    Show bool           `json:"show"`
    Color string        `json:"color"`
}


func GetChampionSupportAlly(c *gin.Context) {
    champion := c.Param("champion")
    championStats := datahandling.GetAdcSupportAlly(champion, BotlaneData)

    totalPlayed := 0.
    data := make([]Data, 0)
    for name, stats := range championStats {
        var color string
        if stats.Winrate >= 55 || stats.Winrate <= 45 {
            color = "#ffffff"
        } else {
            color = "#000000"
        }

        d := Data{
            Name: strings.Title(name),
            Value: []float64{stats.Played, stats.Winrate},
            Label: Label{
                Show: true,
                Color: color,
            },
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
        WinColor: interpolateColor(60, 40, 50, 60, SATURATION),
        NeutralColor: interpolateColor(50, 40, 50, 60, SATURATION),
        LoseColor: interpolateColor(40, 40, 50, 60, SATURATION),
        Data: data,
    }
    c.IndentedJSON(http.StatusOK, options)
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
