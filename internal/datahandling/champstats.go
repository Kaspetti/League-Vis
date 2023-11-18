package datahandling

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
)


type ChampionStats struct {
    Played float64  `json:"played"`
    Wins float64    `json:"wins"`
    Losses float64  `json:"losses"`
    Winrate float64 `json:"winrate"`
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


func GetAdcAllySupport(championName string, botlaneData []BotlaneData) map[string]*ChampionStats {
    championStats := make(map[string]*ChampionStats)

    for _, bd := range botlaneData {
        if !bd.isValid() {
            continue
        }

        var stats *ChampionStats
        var ok bool
        if bd.BottomBlue == championName {
            if stats, ok = championStats[bd.UtilityBlue]; !ok {
                stats = &ChampionStats{}
                championStats[bd.UtilityBlue] = stats
            }
            stats.addGame(100, bd.WinningTeam)
        } else if bd.BottomRed == championName {
            if stats, ok = championStats[bd.UtilityRed]; !ok {
                stats = &ChampionStats{}
                championStats[bd.UtilityRed] = stats
            }
            stats.addGame(200, bd.WinningTeam)
        }
    }

    return championStats
}


func GetAdcOpponentAdc(championName string, botlaneData []BotlaneData) map[string]*ChampionStats {
    championStats := make(map[string]*ChampionStats)

    for _, bd := range botlaneData {
        if !bd.isValid() {
            continue
        }

        var stats *ChampionStats
        var ok bool
        if bd.BottomBlue == championName {
            if stats, ok = championStats[bd.BottomRed]; !ok {
                stats = &ChampionStats{}
                championStats[bd.BottomRed] = stats
            }
            stats.addGame(100, bd.WinningTeam)
        } else if bd.BottomRed == championName {
            if stats, ok = championStats[bd.BottomBlue]; !ok {
                stats = &ChampionStats{}
                championStats[bd.BottomBlue] = stats
            }
            stats.addGame(200, bd.WinningTeam)
        }
    }

    return championStats
}


func (bd *BotlaneData) isValid() bool {
    return  bd.BottomBlue != "" &&
            bd.BottomRed != "" &&
            bd.UtilityBlue != "" &&
            bd.UtilityRed != "" &&
            (bd.WinningTeam == 100 || bd.WinningTeam == 200) &&
            bd.MatchId != ""
}


func (cs *ChampionStats) addGame(playerTeam, winningTeam int) {
    cs.Played++
    if playerTeam == winningTeam {
        cs.Wins++
    } else {
        cs.Losses++
    }

    cs.Winrate = (cs.Wins * 100) / cs.Played
}
