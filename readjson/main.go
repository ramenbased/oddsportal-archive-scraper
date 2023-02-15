package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Year struct {
	Data []struct {
		ID                      int      `json:"id"`
		URL                     string   `json:"url"`
		IsDouble                bool     `json:"is-double"`
		Home                    int      `json:"home"`
		Away                    int      `json:"away"`
		HomeName                string   `json:"home-name"`
		AwayName                string   `json:"away-name"`
		HomeCountryTwoChartName string   `json:"home-country-two-chart-name"`
		AwayCountryTwoChartName string   `json:"away-country-two-chart-name"`
		HomeParticipantID       int      `json:"home-participant-id"`
		AwayParticipantID       int      `json:"away-participant-id"`
		StatusID                int      `json:"status-id"`
		EventStageID            int      `json:"event-stage-id"`
		EventStageName          string   `json:"event-stage-name"`
		TournamentStageID       int      `json:"tournament-stage-id"`
		TournamentStageTypeID   int      `json:"tournament-stage-type-id"`
		TournamentStageGroupID  int      `json:"tournament-stage-group-id"`
		TournamentStageName     string   `json:"tournament-stage-name"`
		SportID                 int      `json:"sport-id"`
		Cols                    string   `json:"cols"`
		CountryID               int      `json:"country-id"`
		CountryName             string   `json:"country-name"`
		CountryTwoChartName     string   `json:"country-two-chart-name"`
		CountryType             string   `json:"country-type"`
		TournamentID            int      `json:"tournament-id"`
		TournamentName          string   `json:"tournament-name"`
		TournamentURL           string   `json:"tournament-url"`
		HomeParticipantImages   []string `json:"home-participant-images"`
		AwayParticipantImages   []string `json:"away-participant-images"`
		SportURLName            string   `json:"sport-url-name"`
		Breadcrumbs             struct {
			Sport struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"sport"`
			Country struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"country"`
			Tournament struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"tournament"`
		} `json:"breadcrumbs"`
		EncodeEventID        string `json:"encodeEventId"`
		ColClassName         string `json:"colClassName"`
		HomeParticipantTypes []int  `json:"homeParticipantTypes"`
		AwayParticipantTypes []int  `json:"awayParticipantTypes"`
		DateStartBase        int    `json:"date-start-base"`
		DateStartTimestamp   int    `json:"date-start-timestamp"`
		Result               string `json:"result"`
		HomeResult           string `json:"homeResult"`
		AwayResult           string `json:"awayResult"`
		HomeWinner           string `json:"home-winner"`
		AwayWinner           string `json:"away-winner"`
		Info                 []struct {
			Name    string `json:"name"`
			SortKey int    `json:"sort_key"`
		} `json:"info"`
		Partialresult   string `json:"partialresult"`
		BookmakersCount int    `json:"bookmakersCount"`
		WinnerPost      int    `json:"winner_post"`
		BettingType     int    `json:"betting_type"`
		Odds            []struct {
			AvgOdds           float64 `json:"avgOdds"`
			BettingTypeID     int     `json:"bettingTypeId"`
			EventID           int     `json:"eventId"`
			MaxOdds           float64 `json:"maxOdds"`
			OutcomeResultID   int     `json:"outcomeResultId"`
			ScopeID           int     `json:"scopeId"`
			OutcomeID         string  `json:"outcomeId"`
			MaxOddsProviderID int     `json:"maxOddsProviderId"`
			Active            bool    `json:"active"`
		} `json:"odds"`
		Name             string `json:"name"`
		ColClassNameTime string `json:"colClassNameTime"`
	} `json:"data"`
}

func Er(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func listEm(y Year) {
	for _, r := range y.Data {
		fmt.Println("---------------------------------------------------------")
		fmt.Printf("%v VS. %v Result: %v \n", r.HomeName+"("+r.HomeWinner+")", r.AwayName+"("+r.AwayWinner+")", r.Result)
		fmt.Println("Date: ", time.Unix(int64(r.DateStartTimestamp), 0))
		fmt.Println(r.Partialresult)
		if len(r.Odds) > 0 {
			fmt.Printf("Odds:\tHome\t%v Average \t %v Max \n\tAway\t%v Average \t %v Max\n", r.Odds[0].AvgOdds, r.Odds[0].MaxOdds, r.Odds[1].AvgOdds, r.Odds[1].MaxOdds)
		}
		fmt.Println("https://oddsportal.com" + r.URL)
	}
}

func avgOddsWinner(y Year) {
	var avgWin float64
	var nWin = 0.0
	var avgLoss float64
	var nLoss = 0.0

	for _, r := range y.Data {
		//WinnerPost is always win for away and no draw
		//home = odds[0], away = odds[1]
		if r.HomeWinner != "draw" && r.HomeWinner != "" {
			if r.HomeWinner == "win" {
				if len(r.Odds) != 0 {
					nWin++
					avgWin += r.Odds[0].AvgOdds
				}
			}
			if r.AwayWinner == "win" {
				if len(r.Odds) != 0 {
					nWin++
					avgWin += r.Odds[1].AvgOdds
				}
			}
			if r.HomeWinner == "lost" {
				if len(r.Odds) != 0 {
					nLoss++
					avgLoss += r.Odds[0].AvgOdds
				}
			}
			if r.AwayWinner == "lost" {
				if len(r.Odds) != 0 {
					nLoss++
					avgLoss += r.Odds[1].AvgOdds
				}
			}
		}

	}
	fmt.Println(avgWin / nWin)
	fmt.Println(avgLoss / nLoss)
	fmt.Println("Check N", nWin, nLoss)
}

func div(y Year) {

	for _, d := range y.Data {
		fmt.Println(d.Info[0].Name)
	}
}

func main() {
	file, err := os.ReadFile("./../../../resources/odds/oddsportal/YearlyMLB/MLB2022.json")
	Er(err)
	var y Year
	if err := json.Unmarshal(file, &y); err != nil {
		Er(err)
	}
	div(y)
	//avgOddsWinner(y)
}
