import requests
from pprint import pprint


def main():
    parse_scores("mlb")
    print()
    parse_scores("nba")
    print()
    parse_scores("nfl")
    print()
    parse_scores("nhl")


def get_json_for_sport(sport):

    if sport == "mlb":
        url = "http://site.api.espn.com/apis/site/v2/sports/baseball/mlb/scoreboard"
    elif sport == "nba":
        url = "http://site.api.espn.com/apis/site/v2/sports/basketball/nba/scoreboard"
    elif sport == "nfl":
        url = "http://site.api.espn.com/apis/site/v2/sports/football/nfl/scoreboard"
    elif sport == "nhl":
        url = "http://site.api.espn.com/apis/site/v2/sports/hockey/nhl/scoreboard"
    else:
        print("An error has occurred.")

    response = requests.get(url)
    if response:
        return response.json()
    else:
        print("An error has occurred.")


def parse_scores(sport):

    data = get_json_for_sport(sport)
    games = []

    for event in data.get("events"):

        game = {}
        game_status = event["status"]["type"]["shortDetail"]
        game.update({"status": game_status})

        for competition in event["competitions"]:
            for competitors in competition["competitors"]:

                homeaway = competitors["homeAway"]
                team = competitors["team"]["abbreviation"]
                score = competitors["score"]
                game.update({homeaway: team + " " + score},)

        games.append(game)

    pprint(games)


if __name__ == "__main__":
    main()

