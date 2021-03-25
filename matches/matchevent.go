package match

import (
	"footballStats/database"
	player "footballStats/players"
	team "footballStats/teams"
	"log"
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

func GetMatchEventById(id int) MatchEvent {
	stmt, err := database.Db.Prepare("SELECT MatchEventID, MatchID, TeamID, PlayerID, EventType, EventMinute, StoppageTime FROM dbo.MatchEvents WHERE MatchEventID = ?")

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(id)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var matchEventFromID MatchEvent

	for rows.Next() {
		var matchEventRow MatchEvent

		err := rows.Scan(&matchEventRow.ID, &matchEventRow.Match.ID, &matchEventRow.Team.ID, &matchEventRow.Player.ID, &matchEventRow.EventType, &matchEventRow.EventMinute, &matchEventRow.StoppageTime)

		if err != nil {
			log.Fatal(err)
		}

		matchEventRow.Match = GetMatchById(matchEventRow.Match.ID)
		matchEventRow.Team = team.GetTeamById(matchEventRow.Team.ID)
		matchEventRow.Player = player.GetPlayerById(matchEventRow.Player.ID)

		matchEventFromID = matchEventRow
	}

	if err != nil {
		log.Fatal(err)
	}

	return matchEventFromID
}

func GetMatchEventsByMatchId(id int) []MatchEvent {
	stmt, err := database.Db.Prepare("SELECT MatchEventID, MatchID, TeamID, PlayerID, EventType, EventMinute, StoppageTime FROM dbo.MatchEvents WHERE MatchID = ?")

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(id)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var matchEvents []MatchEvent

	for rows.Next() {
		var matchEventRow MatchEvent

		err := rows.Scan(&matchEventRow.ID, &matchEventRow.Match.ID, &matchEventRow.Team.ID, &matchEventRow.Player.ID, &matchEventRow.EventType, &matchEventRow.EventMinute, &matchEventRow.StoppageTime)

		if err != nil {
			log.Fatal(err)
		}

		matchEventRow.Match = GetMatchById(matchEventRow.Match.ID)
		matchEventRow.Team = team.GetTeamById(matchEventRow.Team.ID)
		matchEventRow.Player = player.GetPlayerById(matchEventRow.Player.ID)

		matchEvents = append(matchEvents, matchEventRow)
	}

	return matchEvents
}

func (matchEvent MatchEvent) Save() int64 {
	stmt, err := database.Db.Prepare("INSERT INTO dbo.MatchEvents(MatchID, TeamID, PlayerID, EventType, EventMinute, StoppageTime) VALUES(?, ?, ?, ?, ?, ?); SELECT SCOPE_IDENTITY()")

	if err != nil {
		log.Fatal(err)
	}

	var id int64 = 0
	err = stmt.QueryRow(matchEvent.Match.ID, matchEvent.Team.ID, matchEvent.Player.ID, matchEvent.EventType, matchEvent.EventMinute, matchEvent.StoppageTime).Scan(&id)

	defer stmt.Close()

	if err != nil {
		log.Fatal(err)
	}

	return id
}

func UpdateMatchEvent(updatedMatchEvent MatchEvent) MatchEvent {
	stmt, err := database.Db.Prepare("UPDATE dbo.MatchEvents SET MatchID = ?, TeamID = ?, PlayerID = ?, EventType = ?, EventMinute = ?, StoppageTime = ? WHERE MatchEventID = ?")

	if err != nil {
		log.Fatal(err)
	}

	rows, err := stmt.Query(updatedMatchEvent.Match.ID, updatedMatchEvent.Team.ID, updatedMatchEvent.Player.ID, updatedMatchEvent.EventType, updatedMatchEvent.EventMinute, updatedMatchEvent.StoppageTime, updatedMatchEvent.ID)

	if err != nil {
		log.Fatal(err)
	}

	rows.Close()

	return GetMatchEventById(updatedMatchEvent.ID)
}

func DeleteMatchEvent(id int) {
	stmt, err := database.Db.Prepare("DELETE FROM dbo.MatchEvents WHERE MatchEventID = ?")

	if err != nil {
		log.Fatal(err)
	}

	rows, err := stmt.Query(id)

	if err != nil {
		log.Fatal(err)
	}

	rows.Close()
}
