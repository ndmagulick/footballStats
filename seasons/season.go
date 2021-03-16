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
