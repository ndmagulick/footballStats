package league

import (
	"footballStats/database"
	"log"
)

type League struct {
	ID      int
	Name    string
	Country string
	Tier    int
}

func GetAllLeagues() []League {
	stmt, err := database.Db.Prepare("SELECT LeagueID, LeagueName, Country, Tier FROM dbo.Leagues")

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var leagues []League

	for rows.Next() {
		var league League
		err := rows.Scan(&league.ID, &league.Name, &league.Country, &league.Tier)

		if err != nil {
			log.Fatal(err)
		}

		leagues = append(leagues, league)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return leagues
}

func GetLeagueById(id int) League {
	stmt, err := database.Db.Prepare("SELECT LeagueID, LeagueName, Country, Tier FROM dbo.Leagues WHERE LeagueID = ?")

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(id)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var leagueFromID League

	for rows.Next() {
		var leagueRow League
		err := rows.Scan(&leagueRow.ID, &leagueRow.Name, &leagueRow.Country, &leagueRow.Tier)

		if err != nil {
			log.Fatal(err)
		}

		leagueFromID = leagueRow
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return leagueFromID
}

func (league League) Save() int64 {
	stmt, err := database.Db.Prepare("INSERT INTO dbo.Leagues(LeagueName, Country, Tier) VALUES(?, ?, ?); SELECT SCOPE_IDENTITY()")

	if err != nil {
		log.Fatal(err)
	}

	var id int64 = 0
	err = stmt.QueryRow(league.Name, league.Country, league.Tier).Scan(&id)

	defer stmt.Close()

	if err != nil {
		log.Fatal(err)
	}

	log.Print("League inserted!")

	return id
}

func UpdateLeague(updatedLeague League) League {

	stmt, err := database.Db.Prepare("UPDATE dbo.Leagues SET LeagueName = ?, Country = ?, Tier = ? WHERE LeagueID = ?")

	if err != nil {
		log.Fatal(err)
	}

	rows, err := stmt.Query(updatedLeague.Name, updatedLeague.Country, updatedLeague.Tier, updatedLeague.ID)

	if err != nil {
		log.Fatal(err)
	}

	rows.Close()

	return GetLeagueById(updatedLeague.ID)
}

func DeleteLeague(id int) {
	stmt, err := database.Db.Prepare("DELETE FROM dbo.Leagues WHERE LeagueID = ?")

	if err != nil {
		log.Fatal(err)
	}

	rows, err := stmt.Query(id)

	if err != nil {
		log.Fatal(err)
	}

	rows.Close()
}
