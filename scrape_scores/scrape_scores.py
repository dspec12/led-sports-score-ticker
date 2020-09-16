import os
import requests
import boto3


# Config
s3_bucket = os.getenv("S3_BUCKET")


def main():
    all_scores = compile_list_all_scores()
    mlb = "MLB: "
    nba = "NBA: "
    nfl = "NFL: "
    nhl = "NHL: "

    for game in all_scores["mlb"]:
        away = game["away"]
        home = game["home"]
        status = game["status"]
        mlb = mlb + f"{away} at {home} {status}" + " "

    for game in all_scores["nba"]:
        away = game["away"]
        home = game["home"]
        status = game["status"]
        nba = nba + f"{away} at {home} {status}" + " "

    for game in all_scores["nfl"]:
        away = game["away"]
        home = game["home"]
        status = game["status"]
        nfl = nfl + f"{away} at {home} {status}" + " "

    for game in all_scores["nhl"]:
        away = game["away"]
        home = game["home"]
        status = game["status"]
        nhl = nhl + f"{away} at {home} {status}" + " "

    write_to_s3(mlb + nba + nfl + nhl)


def get_json_for_sport(sport):
    if sport == "mlb":
        url = "http://site.api.espn.com/apis/site/v2/sports/baseball/mlb/scoreboard"
    elif sport == "nba":
        url = "http://site.api.espn.com/apis/site/v2/sports/basketball/nba/scoreboard"
    elif sport == "nfl":
        url = "http://site.api.espn.com/apis/site/v2/sports/football/nfl/scoreboard"
    elif sport == "nhl":
        url = "http://site.api.espn.com/apis/site/v2/sports/hockey/nhl/scoreboard"

    try:
        response = requests.get(url)
        return response.json()

    except requests.exceptions.ConnectionError as e:
        print(f"Could not connect to endpoint:")
        print(e)
    except requests.exceptions.HTTPError as e:
        print(f"Http error:")
        print(e)
    except Exception as e:
        print(f"Unknown error:")
        print(type(e))
        print(e)


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
                game.update(
                    {homeaway: team + " " + score},
                )
        games.append(game)

    return games


def compile_list_all_scores():
    mlb = parse_scores("mlb")
    nba = parse_scores("nba")
    nfl = parse_scores("nfl")
    nhl = parse_scores("nhl")

    leagues = {}
    leagues.update({"mlb": mlb, "nba": nba, "nfl": nfl, "nhl": nhl})

    return leagues


def write_to_s3(string):
    client = boto3.client("s3")
    bucket = s3_bucket
    response = client.put_object(
        ACL="public-read", Bucket=bucket, Body=string, Key="scores.txt"
    )


if __name__ == "__main__":
    main()
