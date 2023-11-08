package datahandling

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
)


type ChampionStats struct {
    Played float64
    Wins float64
    Losses float64
    Winrate float64
}


type BotlaneData struct {
    MatchId string
    BottomBlue string
    UtilityBlue string
    BottomRed string
    UtilityRed string
    WinningTeam int
}


func ImportData(path string) ([]BotlaneData, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
    data, err := csvReader.ReadAll()
    if err != nil {
        return nil, err
    }

    botlaneData := make([]BotlaneData, len(data))
    for i, line := range data {
        bd := BotlaneData{
            MatchId: line[0],
            BottomBlue: strings.ToLower(line[1]),
            UtilityBlue: strings.ToLower(line[2]),
            BottomRed: strings.ToLower(line[3]),
            UtilityRed: strings.ToLower(line[4]),
        }
        winningTeam, err := strconv.Atoi(line[5])
        if err != nil {
            return nil, err
        }
        bd.WinningTeam = winningTeam

        botlaneData[i] = bd
    }

    return botlaneData, nil
}


func GetChampionStats(champName string, botlaneData []BotlaneData) map[string]*ChampionStats {
    championStats := make(map[string]*ChampionStats)

    for _, bd := range botlaneData {
        if bd.BottomBlue == champName {
            if stats, ok := championStats[bd.UtilityBlue]; ok {
                stats.Played++

                if bd.WinningTeam == 100 {
                    stats.Wins++
                } else {
                    stats.Losses++
                }

                stats.Winrate = (stats.Wins * 100.0) / stats.Played
            } else {
                stats = &ChampionStats{
                    Played: 1,
                    Winrate: 1,
                }

                if bd.WinningTeam == 100 {
                    stats.Wins = 1
                    stats.Winrate = 100
                } else {
                    stats.Losses = 1
                    stats.Winrate = 0
                }

                championStats[bd.UtilityBlue] = stats
            }
        } else if bd.BottomRed == champName {
            if stats, ok := championStats[bd.UtilityRed]; ok {
                stats.Played++

                if bd.WinningTeam == 200 {
                    stats.Wins++
                } else {
                    stats.Losses++
                }

                stats.Winrate = (stats.Wins * 100) / stats.Played
            } else {
                stats = &ChampionStats{
                    Played: 1,
                    Winrate: 1,
                }

                if bd.WinningTeam == 200 {
                    stats.Wins = 1
                    stats.Winrate = 100
                } else {
                    stats.Losses = 1
                    stats.Winrate = 0
                }

                championStats[bd.UtilityRed] = stats
            }

        }
    }

    return championStats
}
