package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type teamData struct {
	Team struct {
		Name            string `json:"name"`
		DisplayName     string `json:"displayName"`
		StandingSummary string `json:"standingSummary"`
		Record          struct {
			Items []struct {
				Summary string `json:"summary"`
			}
		}
		NextEvent []struct {
			Name         string `json:"name"`
			Competitions []struct {
				Competitors []struct {
					Team struct {
						DisplayName      string `json:"displayName"`
						ShortDisplayName string `json:"shortDisplayName"`
						Abbreviation     string `json:"abbreviation"`
						Nickname         string `json:"nickname"`
					}
					Score struct {
						DisplayValue string `json:"displayValue"`
					}
				}
				Status struct {
					Type struct {
						Name        string `json:"name"`
						State       string `json:"state"`
						Description string `json:"description"`
						Detail      string `json:"detail"`
					}
				}
			}
		}
	}
}

const eaglesURL = "https://site.api.espn.com/apis/site/v2/sports/football/nfl/teams/phi"
const flyersURL = "https://site.api.espn.com/apis/site/v2/sports/hockey/nhl/teams/phi"
const philliesURL = "https://site.api.espn.com/apis/site/v2/sports/baseball/mlb/teams/phi"
const psuURL = "https://site.api.espn.com/apis/site/v2/sports/football/college-football/teams/psu"
const sixersURL = "https://site.api.espn.com/apis/site/v2/sports/basketball/nba/teams/phi"

func getESPNJSON(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	return b
}

func parseJSON(bs []byte) teamData {
	teamInfo := teamData{}
	jsonErr := json.Unmarshal(bs, &teamInfo)
	if jsonErr != nil {
		log.Fatalln(jsonErr)
	}
	return teamInfo
}

// Routes
func handleRequests(p string) {
	http.HandleFunc("/eagles", eagles)
	http.HandleFunc("/flyers", flyers)
	http.HandleFunc("/phillies", phillies)
	http.HandleFunc("/psu", psu)
	http.HandleFunc("/sixers", sixers)
	log.Fatal(http.ListenAndServe(p, nil))
}

// Http Endpoints
func eagles(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: eagles")
	t := getESPNJSON(eaglesURL)
	ts := parseJSON(t)

	teamName := ts.Team.DisplayName
	teamRecord := ts.Team.Record.Items[0].Summary
	nextGame := ts.Team.NextEvent[0].Name
	gameDetail := ts.Team.NextEvent[0].Competitions[0].Status.Type.Detail
	homeTeam := ts.Team.NextEvent[0].Competitions[0].Competitors[0].Team.Nickname
	homeScore := ts.Team.NextEvent[0].Competitions[0].Competitors[0].Score.DisplayValue
	awayTeam := ts.Team.NextEvent[0].Competitions[0].Competitors[1].Team.Nickname
	awayScore := ts.Team.NextEvent[0].Competitions[0].Competitors[1].Score.DisplayValue
	teamStandings := strings.TrimSuffix(ts.Team.StandingSummary, " - East")

	switch gs := ts.Team.NextEvent[0].Competitions[0].Status.Type.State; gs {
	case "pre":
		fmt.Fprintln(w, teamName, "("+teamRecord+")"+":", teamStandings, "-", "Next Game:", nextGame, "-", gameDetail)
	case "post":
		fmt.Fprintln(w, teamName, "("+teamRecord+")"+":", teamStandings, "-", "Last Game:", awayTeam, awayScore, "at", homeTeam, homeScore, "-", gameDetail)
	case "in":
		fmt.Fprintln(w, teamName, "("+teamRecord+")"+":", teamStandings, "-", "Live:", awayTeam, awayScore, "at", homeTeam, homeScore, "-", gameDetail)
	}
}

func flyers(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: flyers")
	t := getESPNJSON(flyersURL)
	ts := parseJSON(t)

	teamName := ts.Team.DisplayName
	teamRecord := ts.Team.Record.Items[0].Summary
	nextGame := ts.Team.NextEvent[0].Name
	gameDetail := ts.Team.NextEvent[0].Competitions[0].Status.Type.Detail
	homeTeam := ts.Team.NextEvent[0].Competitions[0].Competitors[0].Team.Nickname
	homeScore := ts.Team.NextEvent[0].Competitions[0].Competitors[0].Score.DisplayValue
	awayTeam := ts.Team.NextEvent[0].Competitions[0].Competitors[1].Team.Nickname
	awayScore := ts.Team.NextEvent[0].Competitions[0].Competitors[1].Score.DisplayValue
	teamStandings := strings.TrimSuffix(ts.Team.StandingSummary, " - East")

	switch gs := ts.Team.NextEvent[0].Competitions[0].Status.Type.State; gs {
	case "pre":
		fmt.Fprintln(w, teamName, "("+teamRecord+")"+":", teamStandings, "-", "Next Game:", nextGame, "-", gameDetail)
	case "post":
		fmt.Fprintln(w, teamName, "("+teamRecord+")"+":", teamStandings, "-", "Last Game:", awayTeam, awayScore, "at", homeTeam, homeScore, "-", gameDetail)
	case "in":
		fmt.Fprintln(w, teamName, "("+teamRecord+")"+":", teamStandings, "-", "Live:", awayTeam, awayScore, "at", homeTeam, homeScore, "-", gameDetail)
	}
}

func phillies(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: phillies")
	t := getESPNJSON(philliesURL)
	ts := parseJSON(t)

	teamName := ts.Team.DisplayName
	teamRecord := ts.Team.Record.Items[0].Summary
	nextGame := ts.Team.NextEvent[0].Name
	gameDetail := ts.Team.NextEvent[0].Competitions[0].Status.Type.Detail
	homeTeam := ts.Team.NextEvent[0].Competitions[0].Competitors[0].Team.Nickname
	homeScore := ts.Team.NextEvent[0].Competitions[0].Competitors[0].Score.DisplayValue
	awayTeam := ts.Team.NextEvent[0].Competitions[0].Competitors[1].Team.Nickname
	awayScore := ts.Team.NextEvent[0].Competitions[0].Competitors[1].Score.DisplayValue
	teamStandings := strings.TrimSuffix(ts.Team.StandingSummary, " - East")

	switch gs := ts.Team.NextEvent[0].Competitions[0].Status.Type.State; gs {
	case "pre":
		fmt.Fprintln(w, teamName, "("+teamRecord+")"+":", teamStandings, "-", "Next Game:", nextGame, "-", gameDetail)
	case "post":
		fmt.Fprintln(w, teamName, "("+teamRecord+")"+":", teamStandings, "-", "Last Game:", awayTeam, awayScore, "at", homeTeam, homeScore, "-", gameDetail)
	case "in":
		fmt.Fprintln(w, teamName, "("+teamRecord+")"+":", teamStandings, "-", "Live:", awayTeam, awayScore, "at", homeTeam, homeScore, "-", gameDetail)
	}
}

func psu(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: psu")
	t := getESPNJSON(psuURL)
	ts := parseJSON(t)

	teamName := ts.Team.DisplayName
	teamRecord := ts.Team.Record.Items[0].Summary
	nextGame := ts.Team.NextEvent[0].Name
	gameDetail := ts.Team.NextEvent[0].Competitions[0].Status.Type.Detail
	homeTeam := ts.Team.NextEvent[0].Competitions[0].Competitors[0].Team.Nickname
	homeScore := ts.Team.NextEvent[0].Competitions[0].Competitors[0].Score.DisplayValue
	awayTeam := ts.Team.NextEvent[0].Competitions[0].Competitors[1].Team.Nickname
	awayScore := ts.Team.NextEvent[0].Competitions[0].Competitors[1].Score.DisplayValue
	teamStandings := strings.TrimSuffix(ts.Team.StandingSummary, " - East")

	switch gs := ts.Team.NextEvent[0].Competitions[0].Status.Type.State; gs {
	case "pre":
		fmt.Fprintln(w, teamName, "("+teamRecord+")"+":", teamStandings, "-", "Next Game:", nextGame, "-", gameDetail)
	case "post":
		fmt.Fprintln(w, teamName, "("+teamRecord+")"+":", teamStandings, "-", "Last Game:", awayTeam, awayScore, "at", homeTeam, homeScore, "-", gameDetail)
	case "in":
		fmt.Fprintln(w, teamName, "("+teamRecord+")"+":", teamStandings, "-", "Live:", awayTeam, awayScore, "at", homeTeam, homeScore, "-", gameDetail)
	}
}

func sixers(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: sixers")
	t := getESPNJSON(sixersURL)
	ts := parseJSON(t)

	teamName := ts.Team.DisplayName
	teamRecord := ts.Team.Record.Items[0].Summary
	nextGame := ts.Team.NextEvent[0].Name
	gameDetail := ts.Team.NextEvent[0].Competitions[0].Status.Type.Detail
	homeTeam := ts.Team.NextEvent[0].Competitions[0].Competitors[0].Team.Nickname
	homeScore := ts.Team.NextEvent[0].Competitions[0].Competitors[0].Score.DisplayValue
	awayTeam := ts.Team.NextEvent[0].Competitions[0].Competitors[1].Team.Nickname
	awayScore := ts.Team.NextEvent[0].Competitions[0].Competitors[1].Score.DisplayValue
	teamStandings := strings.TrimSuffix(ts.Team.StandingSummary, " - East")

	switch gs := ts.Team.NextEvent[0].Competitions[0].Status.Type.State; gs {
	case "pre":
		fmt.Fprintln(w, teamName, "("+teamRecord+")"+":", teamStandings, "-", "Next Game:", nextGame, "-", gameDetail)
	case "post":
		fmt.Fprintln(w, teamName, "("+teamRecord+")"+":", teamStandings, "-", "Last Game:", awayTeam, awayScore, "at", homeTeam, homeScore, "-", gameDetail)
	case "in":
		fmt.Fprintln(w, teamName, "("+teamRecord+")"+":", teamStandings, "-", "Live:", awayTeam, awayScore, "at", homeTeam, homeScore, "-", gameDetail)
	}
}
