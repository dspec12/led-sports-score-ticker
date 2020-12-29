package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
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
			ShortName    string `json:"shortName"`
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

var teams = map[string]string{
	"eagles":   "https://site.api.espn.com/apis/site/v2/sports/football/nfl/teams/phi",
	"flyers":   "https://site.api.espn.com/apis/site/v2/sports/hockey/nhl/teams/phi",
	"phillies": "https://site.api.espn.com/apis/site/v2/sports/baseball/mlb/teams/phi",
	"psu":      "https://site.api.espn.com/apis/site/v2/sports/football/college-football/teams/psu",
	"sixers":   "https://site.api.espn.com/apis/site/v2/sports/basketball/nba/teams/phi",
}

var c = cache.New(5*time.Minute, 10*time.Minute)

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

func getTeamData(t string) teamData {
	td, found := c.Get(t)
	if found {
		log.Println(t, "found in cache")
		return td.(teamData)
	}
	log.Println(t, "data is stale. Fetching new and caching....")
	teamURL := getESPNJSON(teams[t])
	td = parseJSON(teamURL)
	c.Set(t, td, cache.DefaultExpiration)
	return td.(teamData)
}

func initCache(t map[string]string) {
	for team, url := range t {
		log.Println("Initial load of", team, "data")
		bs := getESPNJSON(url)
		data := parseJSON(bs)
		c.Set(team, data, cache.DefaultExpiration)
	}

}
