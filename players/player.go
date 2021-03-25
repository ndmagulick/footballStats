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

		player.Team = team.GetTeamById(teamID)

		players = append(players, player)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return players
}

func GetPlayerById(id int) Player {
	stmt, err := database.Db.Prepare("SELECT PlayerID, FirstName, LastName, /*DateOfBirth,*/ Height, Nationality, Position, TeamID, Number, Foot FROM dbo.Players WHERE PlayerID = ?")

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(id)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var playerFromID Player

	for rows.Next() {
		var playerRow Player
		teamID := 0

		err := rows.Scan(&playerRow.ID, &playerRow.FirstName, &playerRow.LastName, &playerRow.Height, &playerRow.Nationality, &playerRow.Position, &teamID, &playerRow.Number, &playerRow.Foot)

		if err != nil {
			log.Fatal(err)
		}

		playerRow.Team = team.GetTeamById(teamID)

		playerFromID = playerRow
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return playerFromID
}

func (player Player) Save() int64 {
	stmt, err := database.Db.Prepare("INSERT INTO dbo.Players(FirstName, LastName, Height, Nationality, Position, TeamID, Number, Foot) VALUES(?, ?, ?, ?, ?, ?, ?, ?); SELECT SCOPE_IDENTITY()")

	if err != nil {
		log.Fatal(err)
	}

	var id int64 = 0
	err = stmt.QueryRow(player.FirstName, player.LastName, player.Height, player.Nationality, player.Position, player.Team.ID, player.Number, player.Foot).Scan(&id)

	defer stmt.Close()

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Player inserted!")

	return id
}

func UpdatePlayer(updatedPlayer Player) Player {
	stmt, err := database.Db.Prepare("UPDATE dbo.Players SET Firstname = ?, LastName = ?, Height = ?, Nationality = ?, Position = ?, TeamID = ?, Number = ?, Foot = ? WHERE PlayerID = ?")

	if err != nil {
		log.Fatal(err)
	}

	rows, err := stmt.Query(updatedPlayer.FirstName, updatedPlayer.LastName, updatedPlayer.Height, updatedPlayer.Nationality, updatedPlayer.Position, updatedPlayer.Team.ID,
		updatedPlayer.Number, updatedPlayer.Foot, updatedPlayer.ID)

	if err != nil {
		log.Fatal(err)
	}

	rows.Close()

	return GetPlayerById(updatedPlayer.ID)
}

func DeletePlayer(id int) {
	stmt, err := database.Db.Prepare("DELETE FROM dbo.Players WHERE PlayerID = ?")

	if err != nil {
		log.Fatal(err)
	}

	rows, err := stmt.Query(id)

	if err != nil {
		log.Fatal(err)
	}

	rows.Close()
}
