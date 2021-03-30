// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type DeleteResponse struct {
	Message string `json:"message"`
}

type League struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Tier    int    `json:"tier"`
}

type Match struct {
	ID       int     `json:"id"`
	Season   *Season `json:"season"`
	HomeTeam *Team   `json:"homeTeam"`
	AwayTeam *Team   `json:"awayTeam"`
	MatchDay int     `json:"matchDay"`
}

type MatchEvent struct {
	ID           int     `json:"id"`
	Match        *Match  `json:"match"`
	Team         *Team   `json:"team"`
	Player       *Player `json:"player"`
	EventType    int     `json:"eventType"`
	EventMinute  int     `json:"eventMinute"`
	StoppageTime *int    `json:"stoppageTime"`
}

type MatchStats struct {
	ID                int    `json:"id"`
	Match             *Match `json:"Match"`
	PossessionHome    int    `json:"possessionHome"`
	TotalShotsHome    int    `json:"totalShotsHome"`
	ShotsOnTargetHome int    `json:"shotsOnTargetHome"`
	SavesHome         int    `json:"savesHome"`
	CornersHome       int    `json:"cornersHome"`
	FoulsHome         int    `json:"foulsHome"`
	OffsidesHome      int    `json:"offsidesHome"`
	PossessionAway    int    `json:"possessionAway"`
	TotalShotsAway    int    `json:"totalShotsAway"`
	ShotsOnTargetAway int    `json:"shotsOnTargetAway"`
	SavesAway         int    `json:"savesAway"`
	CornersAway       int    `json:"cornersAway"`
	FoulsAway         int    `json:"foulsAway"`
	OffsidesAway      int    `json:"offsidesAway"`
}

type NewLeague struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	Tier    int    `json:"tier"`
}

type NewMatch struct {
	SeasonID   int `json:"seasonId"`
	HomeTeamID int `json:"homeTeamId"`
	AwayTeamID int `json:"awayTeamId"`
	MatchDay   int `json:"matchDay"`
}

type NewMatchEvent struct {
	MatchID      int  `json:"matchId"`
	TeamID       int  `json:"teamId"`
	PlayerID     int  `json:"playerId"`
	EventType    int  `json:"eventType"`
	EventMinute  int  `json:"eventMinute"`
	StoppageTime *int `json:"stoppageTime"`
}

type NewMatchStats struct {
	MatchID           int `json:"MatchId"`
	PossessionHome    int `json:"possessionHome"`
	TotalShotsHome    int `json:"totalShotsHome"`
	ShotsOnTargetHome int `json:"shotsOnTargetHome"`
	SavesHome         int `json:"savesHome"`
	CornersHome       int `json:"cornersHome"`
	FoulsHome         int `json:"foulsHome"`
	OffsidesHome      int `json:"offsidesHome"`
	PossessionAway    int `json:"possessionAway"`
	TotalShotsAway    int `json:"totalShotsAway"`
	ShotsOnTargetAway int `json:"shotsOnTargetAway"`
	SavesAway         int `json:"savesAway"`
	CornersAway       int `json:"cornersAway"`
	FoulsAway         int `json:"foulsAway"`
	OffsidesAway      int `json:"offsidesAway"`
}

type NewPlayer struct {
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	Height      float64 `json:"height"`
	Nationality string  `json:"nationality"`
	Position    string  `json:"position"`
	TeamID      int     `json:"teamId"`
	Number      int     `json:"number"`
	Foot        string  `json:"foot"`
}

type NewSeason struct {
	LeagueID  int `json:"leagueId"`
	StartYear int `json:"startYear"`
	EndYear   int `json:"endYear"`
}

type NewTeam struct {
	Name         string `json:"name"`
	LeagueID     int    `json:"leagueId"`
	FoundingYear int    `json:"foundingYear"`
}

type Player struct {
	ID          int     `json:"id"`
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	Height      float64 `json:"height"`
	Nationality string  `json:"nationality"`
	Position    string  `json:"position"`
	Team        *Team   `json:"team"`
	Number      int     `json:"number"`
	Foot        string  `json:"foot"`
}

type Season struct {
	ID        int     `json:"id"`
	League    *League `json:"league"`
	StartYear int     `json:"startYear"`
	EndYear   int     `json:"endYear"`
}

type Team struct {
	ID           int     `json:"id"`
	League       *League `json:"league"`
	Name         string  `json:"name"`
	FoundingYear int     `json:"foundingYear"`
}

type UpdatedLeague struct {
	ID      int     `json:"id"`
	Name    *string `json:"name"`
	Country *string `json:"country"`
	Tier    *int    `json:"tier"`
}

type UpdatedMatch struct {
	ID         int  `json:"id"`
	SeasonID   *int `json:"seasonId"`
	HomeTeamID *int `json:"homeTeamId"`
	AwayTeamID *int `json:"awayTeamId"`
	MatchDay   *int `json:"matchDay"`
}

type UpdatedMatchEvent struct {
	ID           int  `json:"id"`
	MatchID      *int `json:"matchId"`
	TeamID       *int `json:"teamId"`
	PlayerID     *int `json:"playerId"`
	EventType    int  `json:"eventType"`
	EventMinute  *int `json:"eventMinute"`
	StoppageTime *int `json:"stoppageTime"`
}

type UpdatedMatchStats struct {
	ID                int  `json:"id"`
	MatchID           *int `json:"MatchId"`
	PossessionHome    *int `json:"possessionHome"`
	TotalShotsHome    *int `json:"totalShotsHome"`
	ShotsOnTargetHome *int `json:"shotsOnTargetHome"`
	SavesHome         *int `json:"savesHome"`
	CornersHome       *int `json:"cornersHome"`
	FoulsHome         *int `json:"foulsHome"`
	OffsidesHome      *int `json:"offsidesHome"`
	PossessionAway    *int `json:"possessionAway"`
	TotalShotsAway    *int `json:"totalShotsAway"`
	ShotsOnTargetAway *int `json:"shotsOnTargetAway"`
	SavesAway         *int `json:"savesAway"`
	CornersAway       *int `json:"cornersAway"`
	FoulsAway         *int `json:"foulsAway"`
	OffsidesAway      *int `json:"offsidesAway"`
}

type UpdatedPlayer struct {
	ID          int      `json:"id"`
	FirstName   *string  `json:"firstName"`
	LastName    *string  `json:"lastName"`
	Height      *float64 `json:"height"`
	Nationality *string  `json:"nationality"`
	Position    *string  `json:"position"`
	TeamID      *int     `json:"teamId"`
	Number      *int     `json:"number"`
	Foot        *string  `json:"foot"`
}

type UpdatedSeason struct {
	ID        int  `json:"id"`
	LeagueID  *int `json:"leagueId"`
	StartYear *int `json:"startYear"`
	EndYear   *int `json:"endYear"`
}

type UpdatedTeam struct {
	ID           int     `json:"id"`
	LeagueID     *int    `json:"leagueId"`
	Name         *string `json:"name"`
	FoundingYear *int    `json:"foundingYear"`
}
