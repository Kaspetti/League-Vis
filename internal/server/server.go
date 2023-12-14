package server

import (
	"fmt"
	"log"

	"github.com/Kaspetti/League-Vis/internal/datahandling"
	"github.com/gin-gonic/gin"
)


var BotlaneData []datahandling.BotlaneData


func RunServer(ip, port string) {
    var err error
    BotlaneData, err = datahandling.ImportData("botlanes.csv")
    if err != nil {
        log.Fatalln(err)
    }

    router := gin.Default()
    router.SetTrustedProxies(nil)

    api := router.Group("/api") 
    {
        api.GET("/champions/:champion/ally/support", GetAllySupport)
        api.GET("/champions/:champion/opponent/adc", GetOpponentADC)
    }

    router.Static("/public", "./public")
    router.StaticFile("/", "./public/index.html")
    router.StaticFile("/riot.txt", "riot.txt")

    //router.NoRoute(func(c *gin.Context) {
    //    c.File("./public/index.html")
    //})

    router.Run(fmt.Sprintf("%s:%s", ip, port))
}
