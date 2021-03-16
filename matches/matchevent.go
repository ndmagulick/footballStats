package match

import (
	player "footballStats/players"
	team "footballStats/teams"
)

type MatchEventType int

const (
	Goal MatchEventType = iota + 1
	Assist
	PenaltyGoal
	PenaltyMiss
	YellowCard
	SecondYellowCard
	RedCard
	SubstitutionOn
	SubstitutionOff
)

type MatchEvent struct {
	ID           int
	Match        Match
	Team         team.Team
	Player       player.Player
	EventType    MatchEventType
	EventMinute  int
	StoppageTime int
}
