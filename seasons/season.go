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

func (season Season) Save() int64 {
	stmt, err := database.Db.Prepare("INSERT INTO dbo.Seasons(LeagueID, StartYear, EndYear) VALUES(?, ?, ?); SELECT SCOPE_IDENTITY")

	if err != nil {
		log.Fatal(err)
	}

	var id int64 = 0
	err = stmt.QueryRow(season.League.ID, season.StartYear, season.EndYear).Scan(&id)

	defer stmt.Close()

	if err != nil {
		log.Fatal(err)
	}

	return id
}

func UpdateSeason(updatedSeason Season) Season {
	stmt, err := database.Db.Prepare("UPDATE dbo.Seasons SET LeagueID = ?, StartYear = ?, EndYear = ? WHERE SeasonID = ?")

	if err != nil {
		log.Fatal(err)
	}

	rows, err := stmt.Query(updatedSeason.League.ID, updatedSeason.StartYear, updatedSeason.EndYear, updatedSeason.ID)

	if err != nil {
		log.Fatal(err)
	}

	rows.Close()

	return GetSeasonById(updatedSeason.ID)
}

func DeleteSeason(id int) {
	stmt, err := database.Db.Prepare("DELETE FROM dbo.Seasons WHERE SeasonID = ?")

	if err != nil {
		log.Fatal(err)
	}

	rows, err := stmt.Query(id)

	if err != nil {
		log.Fatal(err)
	}

	rows.Close()
}
