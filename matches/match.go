package match

import (
	"footballStats/database"
	season "footballStats/seasons"
	team "footballStats/teams"
	"log"
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

func GetMatchById(id int) Match {
	stmt, err := database.Db.Prepare("SELECT MatchID, SeasonID, HomeTeamID, AwayTeamID, MatchDay FROM dbo.Matches WHERE MatchID = ?")

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(id)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var matchFromID Match

	for rows.Next() {
		var matchRow Match
		err := rows.Scan(&matchRow.ID, &matchRow.Season.ID, &matchRow.HomeTeam.ID, &matchRow.AwayTeam.ID, &matchRow.MatchDay)

		if err != nil {
			log.Fatal(err)
		}

		matchRow.Season = season.GetSeasonById(matchRow.Season.ID)
		matchRow.HomeTeam = team.GetTeamById(matchRow.HomeTeam.ID)
		matchRow.AwayTeam = team.GetTeamById(matchRow.AwayTeam.ID)

		matchFromID = matchRow
	}

	return matchFromID
}

func GetMatchesBySeasonId(seasonId int) []Match {
	stmt, err := database.Db.Prepare("SELECT MatchID, SeasonID, HomeTeamID, AwayTeamID, MatchDay FROM dbo.Matches WHERE SeasonID = ?")

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(seasonId)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var matchesFromSeasonId []Match

	for rows.Next() {
		var matchRow Match
		err := rows.Scan(&matchRow.ID, &matchRow.Season.ID, &matchRow.HomeTeam.ID, &matchRow.AwayTeam.ID, &matchRow.MatchDay)

		if err != nil {
			log.Fatal(err)
		}

		matchRow.Season = season.GetSeasonById(matchRow.Season.ID)
		matchRow.HomeTeam = team.GetTeamById(matchRow.HomeTeam.ID)
		matchRow.AwayTeam = team.GetTeamById(matchRow.AwayTeam.ID)

		matchesFromSeasonId = append(matchesFromSeasonId, matchRow)
	}

	return matchesFromSeasonId
}

func GetAllMatchesByTeamId(teamId int) []Match {
	stmt, err := database.Db.Prepare("SELECT MatchID, SeasonID, HomeTeamID, AwayTeamID, MatchDay FROM dbo.Matches WHERE HomeTeamID = ? OR AwayTeamID = ?")

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(teamId, teamId)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var matchesFromTeamId []Match

	for rows.Next() {
		var matchRow Match
		err := rows.Scan(&matchRow.ID, &matchRow.Season.ID, &matchRow.HomeTeam.ID, &matchRow.AwayTeam.ID, &matchRow.MatchDay)

		if err != nil {
			log.Fatal(err)
		}

		matchRow.Season = season.GetSeasonById(matchRow.Season.ID)
		matchRow.HomeTeam = team.GetTeamById(matchRow.HomeTeam.ID)
		matchRow.AwayTeam = team.GetTeamById(matchRow.AwayTeam.ID)

		matchesFromTeamId = append(matchesFromTeamId, matchRow)
	}

	return matchesFromTeamId
}

func GetMatchesFromTeamAndSeasonId(teamId int, seasonId int) []Match {
	stmt, err := database.Db.Prepare("SELECT MatchID, SeasonID, HomeTeamID, AwayTeamID, MatchDay FROM dbo.Matches WHERE (HomeTeamID = ? OR AwayTeamID = ?) AND SeasonID = ?")

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(teamId, teamId, seasonId)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var matchesFromTeamId []Match

	for rows.Next() {
		var matchRow Match
		err := rows.Scan(&matchRow.ID, &matchRow.Season.ID, &matchRow.HomeTeam.ID, &matchRow.AwayTeam.ID, &matchRow.MatchDay)

		if err != nil {
			log.Fatal(err)
		}

		matchRow.Season = season.GetSeasonById(matchRow.Season.ID)
		matchRow.HomeTeam = team.GetTeamById(matchRow.HomeTeam.ID)
		matchRow.AwayTeam = team.GetTeamById(matchRow.AwayTeam.ID)

		matchesFromTeamId = append(matchesFromTeamId, matchRow)
	}

	return matchesFromTeamId
}

func (match Match) Save() int64 {
	stmt, err := database.Db.Prepare("INSERT INTO dbo.Matches(SeasonID, HomeTeamID, AwayTeamID, MatchDay)VALUES(?, ?, ?, ?)")

	if err != nil {
		log.Fatal(err)
	}

	var id int64 = 0
	err = stmt.QueryRow(match.Season.ID, match.HomeTeam.ID, match.AwayTeam.ID, match.MatchDay).Scan(&id)

	defer stmt.Close()

	if err != nil {
		log.Fatal(err)
	}

	return id
}

func UpdateMatch(updatedMatch Match) Match {
	stmt, err := database.Db.Prepare("UPDATE dbo.Matches SET SeasonID = ?, HomeTeamID = ?, AwayTeamID = ?, MatchDay = ? WHERE MatchID = ?")

	if err != nil {
		log.Fatal(err)
	}

	rows, err := stmt.Query(updatedMatch.Season.ID, updatedMatch.HomeTeam.ID, updatedMatch.AwayTeam.ID, updatedMatch.MatchDay, updatedMatch.ID)

	if err != nil {
		log.Fatal(err)
	}

	rows.Close()

	return GetMatchById(updatedMatch.ID)
}

func DeleteMatch(id int) {
	stmt, err := database.Db.Prepare("DELETE FROM dbo.Matches WHERE MatchID = ?")

	if err != nil {
		log.Fatal(err)
	}

	rows, err := stmt.Query(id)

	if err != nil {
		log.Fatal(err)
	}

	rows.Close()
}
