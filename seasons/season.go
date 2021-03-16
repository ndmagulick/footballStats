package season

import league "footballStats/leagues"

type Season struct {
	ID        int
	League    league.League
	StartYear int
	EndYear   int
}
