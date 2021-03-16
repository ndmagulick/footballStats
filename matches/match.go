package match

import (
	season "footballStats/seasons"
	team "footballStats/teams"
)

type Match struct {
	ID          int
	Season      season.Season
	HomeTeam    team.Team
	AwayTeam    team.Team
	MatchDay    int
	MatchEvents []MatchEvent
	MatchStats  MatchStats
}
