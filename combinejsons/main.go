package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type ResultsRaw struct {
	S int `json:"-"`
	D struct {
		Total   int `json:"total"`
		OnePage int `json:"onePage"`
		Page    int `json:"page"`
		Rows    []struct {
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
			Partialresult string `json:"partialresult"`
			Stream        struct {
				Num2   int `json:"2"`
				Num16  int `json:"16"`
				Num45  int `json:"45"`
				Num49  int `json:"49"`
				Num141 int `json:"141"`
				Num411 int `json:"411"`
			} `json:"-"`
			BookmakersCount int `json:"bookmakersCount"`
			WinnerPost      int `json:"winner_post"`
			BettingType     int `json:"betting_type"`
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
		} `json:"rows"`
		PaginationView string `json:"-"`
	} `json:"d"`
	Refresh int `json:"-"`
}

//STREAM IGNORED
type SingleRow struct {
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
	Partialresult string `json:"partialresult"`
	Stream        struct {
		Num2   int `json:"2"`
		Num16  int `json:"16"`
		Num45  int `json:"45"`
		Num49  int `json:"49"`
		Num141 int `json:"141"`
		Num411 int `json:"411"`
	} `json:"-"`
	BookmakersCount int `json:"bookmakersCount"`
	WinnerPost      int `json:"winner_post"`
	BettingType     int `json:"betting_type"`
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
}

func Er(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

type ResultsCompiled struct {
	Rows []SingleRow `json:"data"`
}

func main() {
	var sourceFolder string
	var saveAs string

	flag.StringVar(&sourceFolder, "f", "../results/2022", "source folder with jsons")
	flag.StringVar(&saveAs, "s", "MLB2022", "save as")
	flag.Parse()

	var ResultsCompiled_ ResultsCompiled
	countRows := 0

	// To do: Test Paths on Windows..

	path := filepath.FromSlash(sourceFolder)
	dir, err := os.ReadDir(path)
	Er(err)
	for _, f := range dir {
		if f.Type().IsRegular() {
			if strings.Contains(f.Name(), ".json") {
				// Checks integrity of scraped jsons
				// Checks last two runes of Filename agsainst JSON.Page
				fn := strings.TrimSuffix(f.Name(), ".json")
				lfn := len(fn)
				checkpage, err := strconv.Atoi(fn[lfn-2:])
				Er(err)
				abs, err := filepath.Abs(path)
				Er(err)
				fp := filepath.Join(abs, f.Name())
				file, err := os.ReadFile(fp)
				Er(err)
				fmt.Printf("CHECKING AND MERGING: %v \n", fp)
				var ResultsRaw_ ResultsRaw
				if err := json.Unmarshal(file, &ResultsRaw_); err != nil {
					fmt.Println(err)
				}
				if ResultsRaw_.D.Page != checkpage {
					fmt.Println(ResultsRaw_.D.Page, checkpage)
					os.Exit(1)
				}
				for _, row := range ResultsRaw_.D.Rows {
					countRows += 1
					ResultsCompiled_.Rows = append(ResultsCompiled_.Rows, row)
				}
			}
		}
	}
	output, err := json.Marshal(ResultsCompiled_)
	Er(err)
	os.WriteFile(saveAs+".json", output, 0644)
	fmt.Printf("SUCCESS! Compiled %v rows of games", countRows)
}
