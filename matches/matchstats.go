package match

import (
	"footballStats/database"
	"log"
)

type MatchStats struct {
	ID                int
	MatchID           int
	PossessionHome    int
	TotalShotsHome    int
	ShotsOnTargetHome int
	SavesHome         int
	CornersHome       int
	FoulsHome         int
	OffsidesHome      int
	PossessionAway    int
	TotalShotsAway    int
	ShotsOnTargetAway int
	SavesAway         int
	CornersAway       int
	FoulsAway         int
	OffsidesAway      int
}

func GetMatchStatsById(id int) MatchStats {
	stmt, err := database.Db.Prepare("SELECT MatchStatsID, MatchID, PossessionHome, TotalShotsHome, ShotsOnTargetHome, SavesHome, CornersHome, FoulsHome, OffsidesHome, PossessionAway, TotalShotsAway, ShotsOnTargetAway, SavesAway, CornersAway, FoulsAway, OffsidesAway FROM dbo.MatchStats WHERE MatchStatsID = ?")

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(id)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var matchStatsFromID MatchStats

	for rows.Next() {
		var matchStatsRow MatchStats

		err := rows.Scan(&matchStatsRow.ID, &matchStatsRow.MatchID, &matchStatsRow.PossessionHome, &matchStatsRow.TotalShotsHome, &matchStatsRow.ShotsOnTargetHome, &matchStatsRow.SavesHome, &matchStatsRow.CornersHome, &matchStatsRow.FoulsHome, &matchStatsRow.OffsidesHome, &matchStatsRow.PossessionAway, &matchStatsRow.TotalShotsAway, &matchStatsRow.ShotsOnTargetAway, &matchStatsRow.SavesAway, &matchStatsRow.CornersAway, &matchStatsRow.FoulsAway, &matchStatsRow.OffsidesAway)

		if err != nil {
			log.Fatal(err)
		}

		matchStatsFromID = matchStatsRow
	}

	return matchStatsFromID
}

func GetMatchStatsByMatchId(id int) MatchStats {
	stmt, err := database.Db.Prepare("SELECT MatchStatsID, MatchID, PossessionHome, TotalShotsHome, ShotsOnTargetHome, SavesHome, CornersHome, FoulsHome, OffsidesHome, PossessionAway, TotalShotsAway, ShotsOnTargetAway, SavesAway, CornersAway, FoulsAway, OffsidesAway FROM dbo.MatchStats WHERE MatchID = ?")

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(id)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var matchStatsFromID MatchStats

	for rows.Next() {
		var matchStatsRow MatchStats

		err := rows.Scan(&matchStatsRow.ID, &matchStatsRow.MatchID, &matchStatsRow.PossessionHome, &matchStatsRow.TotalShotsHome, &matchStatsRow.ShotsOnTargetHome, &matchStatsRow.SavesHome, &matchStatsRow.CornersHome, &matchStatsRow.FoulsHome, &matchStatsRow.OffsidesHome, &matchStatsRow.PossessionAway, &matchStatsRow.TotalShotsAway, &matchStatsRow.ShotsOnTargetAway, &matchStatsRow.SavesAway, &matchStatsRow.CornersAway, &matchStatsRow.FoulsAway, &matchStatsRow.OffsidesAway)

		if err != nil {
			log.Fatal(err)
		}

		matchStatsFromID = matchStatsRow
	}

	return matchStatsFromID
}

func (matchStats MatchStats) Save() int64 {
	stmt, err := database.Db.Prepare("INSERT INTO dbo.MatchEvents(MatchID, PossessionHome, TotalShotsHome, ShotsOnTargetHome, SavesHome, CornersHome, FoulsHome, OffsidesHome, PossessionAway, TotalShotsAway, ShotsOnTargetAway, SavesAway, CornersAway, FoulsAway, OffsidesAway) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?); SELECT SCOPE_IDENTITY()")

	if err != nil {
		log.Fatal(err)
	}

	var id int64 = 0
	err = stmt.QueryRow(matchStats.MatchID, matchStats.PossessionHome, matchStats.TotalShotsHome, matchStats.ShotsOnTargetHome, matchStats.SavesHome, matchStats.CornersHome, matchStats.FoulsHome, matchStats.OffsidesHome, matchStats.PossessionAway, matchStats.TotalShotsAway, matchStats.ShotsOnTargetAway, matchStats.SavesAway, matchStats.CornersAway, matchStats.FoulsAway, matchStats.OffsidesAway).Scan(&id)

	defer stmt.Close()

	if err != nil {
		log.Fatal(err)
	}

	return id
}

func UpdateMatchStats(updatedMatchStats MatchStats) MatchStats {
	stmt, err := database.Db.Prepare("UPDATE dbo.MatchStats SET MatchID = ?, PossessionHome = ?, TotalShotsHome = ?, ShotsOnTargetHome = ?, SavesHome = ?, CornersHome = ?, FoulsHome = ?, OffsidesHome = ?, PossessionAway = ?, TotalShotsAway = ?, ShotsOnTargetAway = ?, SavesAway = ?, CornersAway = ?, FoulsAway = ?, OffsidesAway = ? WHERE MatchStatsID = ?")

	if err != nil {
		log.Fatal(err)
	}

	rows, err := stmt.Query(updatedMatchStats.MatchID, updatedMatchStats.PossessionHome, updatedMatchStats.TotalShotsHome, updatedMatchStats.ShotsOnTargetHome, updatedMatchStats.SavesHome, updatedMatchStats.CornersHome, updatedMatchStats.FoulsHome, updatedMatchStats.OffsidesHome, updatedMatchStats.PossessionAway, updatedMatchStats.TotalShotsAway, updatedMatchStats.ShotsOnTargetAway, updatedMatchStats.SavesAway, updatedMatchStats.CornersAway, updatedMatchStats.FoulsAway, updatedMatchStats.OffsidesAway, updatedMatchStats.ID)

	if err != nil {
		log.Fatal(err)
	}

	rows.Close()

	return GetMatchStatsById(updatedMatchStats.ID)
}

func DeleteMatchStats(id int) {
	stmt, err := database.Db.Prepare("DELETE FROM dbo.MatchStats WHERE MatchStatsID = ?")

	if err != nil {
		log.Fatal(err)
	}

	rows, err := stmt.Query(id)

	if err != nil {
		log.Fatal(err)
	}

	rows.Close()
}
