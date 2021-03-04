package league

import (
	"footballStats/database"
	"log"
	"strconv"
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
	stmt, err := database.Db.Prepare("SELECT LeagueID, LeagueName, Country, Tier FROM dbo.Leagues WHERE LeagueID = " + strconv.Itoa(id))

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var leagueFromID League

	for rows.Next() {
		var league League
		err := rows.Scan(&league.ID, &league.Name, &league.Country, &league.Tier)

		if err != nil {
			log.Fatal(err)
		}

		leagueFromID = league
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

	if err != nil {
		log.Fatal(err)
	}

	log.Print("League inserted!")

	return id
}
