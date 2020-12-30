package main

import (
	"crypto/subtle"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Routes
func handleRequests(p string, u string, pw string) {
	http.HandleFunc("/eagles", basicAuth(eagles, u, pw))
	http.HandleFunc("/flyers", basicAuth(flyers, u, pw))
	http.HandleFunc("/phillies", basicAuth(phillies, u, pw))
	http.HandleFunc("/psu", basicAuth(psu, u, pw))
	http.HandleFunc("/sixers", basicAuth(sixers, u, pw))
	log.Fatal(http.ListenAndServe(p, nil))
}

// Auth Middleware
func basicAuth(handler http.HandlerFunc, username string, password string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()

		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(password)) != 1 {
			realm := "Must provide a username and password"
			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			log.Println("Unauthorized hit:", r)
			return
		}
		handler(w, r)
	}
}

// Http Endpoints
func eagles(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: eagles")
	ts := getTeamData("eagles")
	teamName := ts.Team.Name

	if len(ts.Team.Record.Items) == 0 || len(ts.Team.NextEvent) == 0 {
		fmt.Fprintln(w, teamName+":", "No scheduled games")
	} else {
		teamRecord := ts.Team.Record.Items[0].Summary
		nextGame := ts.Team.NextEvent[0].ShortName
		gameDetail := ts.Team.NextEvent[0].Competitions[0].Status.Type.Detail
		homeTeam := ts.Team.NextEvent[0].Competitions[0].Competitors[0].Team.Abbreviation
		homeScore := ts.Team.NextEvent[0].Competitions[0].Competitors[0].Score.DisplayValue
		awayTeam := ts.Team.NextEvent[0].Competitions[0].Competitors[1].Team.Abbreviation
		awayScore := ts.Team.NextEvent[0].Competitions[0].Competitors[1].Score.DisplayValue
		teamStandings := strings.TrimSuffix(ts.Team.StandingSummary, " - East")

		switch gs := ts.Team.NextEvent[0].Competitions[0].Status.Type.State; gs {
		case "pre":
			fmt.Fprintln(w, teamName, "("+teamRecord+")"+":", teamStandings, "-", "Next Game:", nextGame, "-", gameDetail)
		case "post":
			fmt.Fprintln(w, teamName, "("+teamRecord+")"+":", teamStandings, "-", "Last Game:", awayTeam, awayScore, "@", homeTeam, homeScore, "-", gameDetail)
		case "in":
			fmt.Fprintln(w, teamName, "("+teamRecord+")"+":", teamStandings, "-", "Live:", awayTeam, awayScore, "@", homeTeam, homeScore, "-", gameDetail)
		}
	}
}

func flyers(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: flyers")
	ts := getTeamData("flyers")
	teamName := ts.Team.Name

	if len(ts.Team.Record.Items) == 0 || len(ts.Team.NextEvent) == 0 {
		fmt.Fprintln(w, teamName+":", "No scheduled games")
	} else {
		teamRecord := ts.Team.Record.Items[0].Summary
		nextGame := ts.Team.NextEvent[0].ShortName
		gameDetail := ts.Team.NextEvent[0].Competitions[0].Status.Type.Detail
		homeTeam := ts.Team.NextEvent[0].Competitions[0].Competitors[0].Team.Abbreviation
		homeScore := ts.Team.NextEvent[0].Competitions[0].Competitors[0].Score.DisplayValue
		awayTeam := ts.Team.NextEvent[0].Competitions[0].Competitors[1].Team.Abbreviation
		awayScore := ts.Team.NextEvent[0].Competitions[0].Competitors[1].Score.DisplayValue
		teamStandings := strings.TrimSuffix(ts.Team.StandingSummary, " - East")

		switch gs := ts.Team.NextEvent[0].Competitions[0].Status.Type.State; gs {
		case "pre":
			fmt.Fprintln(w, teamName, "("+teamRecord+")"+":", teamStandings, "-", "Next Game:", nextGame, "-", gameDetail)
		case "post":
			fmt.Fprintln(w, teamName, "("+teamRecord+")"+":", teamStandings, "-", "Last Game:", awayTeam, awayScore, "@", homeTeam, homeScore, "-", gameDetail)
		case "in":
			fmt.Fprintln(w, teamName, "("+teamRecord+")"+":", teamStandings, "-", "Live:", awayTeam, awayScore, "@", homeTeam, homeScore, "-", gameDetail)
		}
	}
}

func phillies(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: phillies")
	ts := getTeamData("phillies")
	teamName := ts.Team.Name

	if len(ts.Team.Record.Items) == 0 || len(ts.Team.NextEvent) == 0 {
		fmt.Fprintln(w, teamName+":", "No scheduled games")
	} else {
		teamRecord := ts.Team.Record.Items[0].Summary
		nextGame := ts.Team.NextEvent[0].ShortName
		gameDetail := ts.Team.NextEvent[0].Competitions[0].Status.Type.Detail
		homeTeam := ts.Team.NextEvent[0].Competitions[0].Competitors[0].Team.Abbreviation
		homeScore := ts.Team.NextEvent[0].Competitions[0].Competitors[0].Score.DisplayValue
		awayTeam := ts.Team.NextEvent[0].Competitions[0].Competitors[1].Team.Abbreviation
		awayScore := ts.Team.NextEvent[0].Competitions[0].Competitors[1].Score.DisplayValue
		teamStandings := strings.TrimSuffix(ts.Team.StandingSummary, " - East")

		switch gs := ts.Team.NextEvent[0].Competitions[0].Status.Type.State; gs {
		case "pre":
			fmt.Fprintln(w, teamName, "("+teamRecord+")"+":", teamStandings, "-", "Next Game:", nextGame, "-", gameDetail)
		case "post":
			fmt.Fprintln(w, teamName, "("+teamRecord+")"+":", teamStandings, "-", "Last Game:", awayTeam, awayScore, "@", homeTeam, homeScore, "-", gameDetail)
		case "in":
			fmt.Fprintln(w, teamName, "("+teamRecord+")"+":", teamStandings, "-", "Live:", awayTeam, awayScore, "@", homeTeam, homeScore, "-", gameDetail)
		}
	}
}

func psu(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: psu")
	ts := getTeamData("psu")
	teamName := ts.Team.Name

	if len(ts.Team.Record.Items) == 0 || len(ts.Team.NextEvent) == 0 {
		fmt.Fprintln(w, teamName+":", "No scheduled games")
	} else {
		teamRecord := ts.Team.Record.Items[0].Summary
		nextGame := ts.Team.NextEvent[0].ShortName
		gameDetail := ts.Team.NextEvent[0].Competitions[0].Status.Type.Detail
		homeTeam := ts.Team.NextEvent[0].Competitions[0].Competitors[0].Team.Abbreviation
		homeScore := ts.Team.NextEvent[0].Competitions[0].Competitors[0].Score.DisplayValue
		awayTeam := ts.Team.NextEvent[0].Competitions[0].Competitors[1].Team.Abbreviation
		awayScore := ts.Team.NextEvent[0].Competitions[0].Competitors[1].Score.DisplayValue
		teamStandings := strings.TrimSuffix(ts.Team.StandingSummary, " - East")

		switch gs := ts.Team.NextEvent[0].Competitions[0].Status.Type.State; gs {
		case "pre":
			fmt.Fprintln(w, teamName, "("+teamRecord+")"+":", teamStandings, "-", "Next Game:", nextGame, "-", gameDetail)
		case "post":
			fmt.Fprintln(w, teamName, "("+teamRecord+")"+":", teamStandings, "-", "Last Game:", awayTeam, awayScore, "@", homeTeam, homeScore, "-", gameDetail)
		case "in":
			fmt.Fprintln(w, teamName, "("+teamRecord+")"+":", teamStandings, "-", "Live:", awayTeam, awayScore, "@", homeTeam, homeScore, "-", gameDetail)
		}
	}
}

func sixers(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: sixers")
	ts := getTeamData("sixers")
	teamName := ts.Team.Name

	if len(ts.Team.Record.Items) == 0 || len(ts.Team.NextEvent) == 0 {
		fmt.Fprintln(w, teamName+":", "No scheduled games")
	} else {
		teamRecord := ts.Team.Record.Items[0].Summary
		nextGame := ts.Team.NextEvent[0].ShortName
		gameDetail := ts.Team.NextEvent[0].Competitions[0].Status.Type.Detail
		homeTeam := ts.Team.NextEvent[0].Competitions[0].Competitors[0].Team.Abbreviation
		homeScore := ts.Team.NextEvent[0].Competitions[0].Competitors[0].Score.DisplayValue
		awayTeam := ts.Team.NextEvent[0].Competitions[0].Competitors[1].Team.Abbreviation
		awayScore := ts.Team.NextEvent[0].Competitions[0].Competitors[1].Score.DisplayValue
		teamStandings := strings.TrimSuffix(ts.Team.StandingSummary, " - East")

		switch gs := ts.Team.NextEvent[0].Competitions[0].Status.Type.State; gs {
		case "pre":
			fmt.Fprintln(w, teamName, "("+teamRecord+")"+":", teamStandings, "-", "Next Game:", nextGame, "-", gameDetail)
		case "post":
			fmt.Fprintln(w, teamName, "("+teamRecord+")"+":", teamStandings, "-", "Last Game:", awayTeam, awayScore, "@", homeTeam, homeScore, "-", gameDetail)
		case "in":
			fmt.Fprintln(w, teamName, "("+teamRecord+")"+":", teamStandings, "-", "Live:", awayTeam, awayScore, "@", homeTeam, homeScore, "-", gameDetail)
		}
	}
}
