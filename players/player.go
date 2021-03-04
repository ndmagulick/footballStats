package player

import (
	"footballStats/database"
	team "footballStats/teams"
	"log"
)

type Player struct {
	ID          int
	FirstName   string
	LastName    string
	Height      float64
	Nationality string
	Position    string
	Team        team.Team
	Number      int
	Foot        string
}

func GetAllPlayers() []Player {
	stmt, err := database.Db.Prepare("SELECT PlayerID, FirstName, LastName, /*DateOfBirth,*/ Height, Nationality, Position, TeamID, Number, Foot FROM dbo.Players")

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var players []Player

	for rows.Next() {
		var player Player
		teamID := 0

		err := rows.Scan(&player.ID, &player.FirstName, &player.LastName, &player.Height, &player.Nationality, &player.Position, &teamID, &player.Number, &player.Foot)

		if err != nil {
			log.Fatal(err)
		}

		player.Team = team.GetTeamFromId(teamID)

		players = append(players, player)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return players
}

func (player Player) Save() int64 {
	stmt, err := database.Db.Prepare("INSERT INTO dbo.Players(FirstName, LastName, Height, Nationality, Position, TeamID, Number, Foot) VALUES(?, ?, ?, ?, ?, ?, ?, ?); SELECT SCOPE_IDENTITY()")

	if err != nil {
		log.Fatal(err)
	}

	var id int64 = 0
	err = stmt.QueryRow(player.FirstName, player.LastName, player.Height, player.Nationality, player.Position, player.Team.ID, player.Number, player.Foot).Scan(&id)

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Player inserted!")

	return id
}
