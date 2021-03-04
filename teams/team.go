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

	var teamFromID Team

	for rows.Next() {
		var team Team
		leagueID := 0

		err := rows.Scan(&team.ID, &leagueID, &team.Name, &team.FoundingYear)

		if err != nil {
			log.Fatal(err)
		}

		team.League = league.GetLeagueById(leagueID)

		teamFromID = team
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
