package team

import (
	"footballStats/database"
	league "footballStats/leagues"
	"log"
)

type Team struct {
	ID           int
	League       league.League
	Name         string
	FoundingYear int
}

func GetAllTeams() []Team {
	stmt, err := database.Db.Prepare("SELECT TeamID, LeagueID, TeamName, FoundingYear FROM dbo.Teams")

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var teams []Team

	for rows.Next() {
		var team Team
		leagueID := 0

		err := rows.Scan(&team.ID, &leagueID, &team.Name, &team.FoundingYear)

		if err != nil {
			log.Fatal(err)
		}

		team.League = league.GetLeagueById(leagueID)

		teams = append(teams, team)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return teams
}

func GetTeamById(id int) Team {
	stmt, err := database.Db.Prepare("SELECT TeamID, LeagueID, TeamName, FoundingYear FROM dbo.Teams WHERE TeamId = ?")

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(id)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var teamFromID Team

	for rows.Next() {
		var teamRow Team
		leagueID := 0

		err := rows.Scan(&teamRow.ID, &leagueID, &teamRow.Name, &teamRow.FoundingYear)

		if err != nil {
			log.Fatal(err)
		}

		teamRow.League = league.GetLeagueById(leagueID)

		teamFromID = teamRow
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return teamFromID
}

func (team Team) Save() int64 {
	stmt, err := database.Db.Prepare("INSERT INTO dbo.Teams(LeagueId, TeamName, FoundingYear) VALUES(?, ?, ?); SELECT SCOPE_IDENTITY()")

	if err != nil {
		log.Fatal(err)
	}

	var id int64 = 0
	err = stmt.QueryRow(team.League.ID, team.Name, team.FoundingYear).Scan(&id)

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Team inserted!")

	return id
}

func UpdateTeam(updatedTeam Team) Team {
	stmt, err := database.Db.Prepare("UPDATE dbo.Teams SET LeagueID = ?, TeamName = ?, FoundingYear = ? WHERE TeamID = ?")

	if err != nil {
		log.Fatal(err)
	}

	rows, err := stmt.Query(updatedTeam.League.ID, updatedTeam.Name, updatedTeam.FoundingYear, updatedTeam.ID)

	if err != nil {
		log.Fatal(err)
	}

	rows.Close()

	return GetTeamById(updatedTeam.ID)
}

func DeleteTeam(id int) {
	stmt, err := database.Db.Prepare("DELETE FROM dbo.Teams WHERE TeamID = ?")

	if err != nil {
		log.Fatal(err)
	}

	rows, err := stmt.Query(id)

	if err != nil {
		log.Fatal(err)
	}

	rows.Close()
}
