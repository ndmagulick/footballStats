package season

import (
	"footballStats/database"
	league "footballStats/leagues"
	"log"
)

type Season struct {
	ID        int
	League    league.League
	StartYear int
	EndYear   int
	//Matches   []match.Match
}

func GetAllSeasonsByLeagueId(leagueId int) []Season {
	stmt, err := database.Db.Prepare("SELECT SeasonID, StartYear, EndYear FROM dbo.Seasons WHERE LeagueID = ?")

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(leagueId)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var seasonsByLeagueId []Season
	league := league.GetLeagueById(leagueId)

	for rows.Next() {
		var season Season

		err := rows.Scan(&season.ID, &season.StartYear, &season.EndYear)

		if err != nil {
			log.Fatal(err)
		}

		season.League = league
		seasonsByLeagueId = append(seasonsByLeagueId, season)
	}

	return seasonsByLeagueId
}

func GetSeasonByLeagueAndStartYear(leagueId int, startYear int) Season {
	stmt, err := database.Db.Prepare("SELECT SeasonID, StartYear, EndYear FROM dbo.Seasons WHERE LeagueID = ? AND StartYear = ?")

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(leagueId, startYear)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var seasonsByLeagueIdAndEndYear Season

	for rows.Next() {
		var seasonRow Season

		err := rows.Scan(&seasonRow.ID, &seasonRow.StartYear, &seasonRow.EndYear)

		if err != nil {
			log.Fatal(err)
		}

		seasonRow.League = league.GetLeagueById(leagueId)
		seasonsByLeagueIdAndEndYear = seasonRow
	}

	return seasonsByLeagueIdAndEndYear
}

func GetSeasonById(seasonId int) Season {
	stmt, err := database.Db.Prepare("SELECT SeasonID, LeagueID, StartYear, EndYear FROM dbo.Seasons WHERE SeasonID = ?")

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(seasonId)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var seasonById Season

	for rows.Next() {
		var seasonRow Season

		err := rows.Scan(&seasonRow.ID, &seasonRow.League.ID, &seasonRow.StartYear, &seasonRow.EndYear)

		if err != nil {
			log.Fatal(err)
		}

		seasonRow.League = league.GetLeagueById(seasonRow.League.ID)
		seasonById = seasonRow
	}

	return seasonById
}

//get season by id
//save
//update
//delete
