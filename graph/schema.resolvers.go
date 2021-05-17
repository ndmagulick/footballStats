package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"footballStats/graph/generated"
	"footballStats/graph/model"
	league "footballStats/leagues"
	match "footballStats/matches"
	player "footballStats/players"
	season "footballStats/seasons"
	team "footballStats/teams"
)

func (r *mutationResolver) CreateLeague(ctx context.Context, input model.NewLeague) (*model.League, error) {
	var newLeague league.League
	newLeague.Name = input.Name
	newLeague.Country = input.Country
	newLeague.Tier = input.Tier

	newLeagueID := newLeague.Save()

	return &model.League{
			ID:      int(newLeagueID),
			Name:    newLeague.Name,
			Country: newLeague.Country,
			Tier:    newLeague.Tier},
		nil
}

func (r *mutationResolver) UpdateLeague(ctx context.Context, input model.UpdatedLeague) (*model.League, error) {
	leagueToUpdate := league.GetLeagueById(input.ID)

	if input.Name != nil && len(*input.Name) > 0 {
		leagueToUpdate.Name = *input.Name
	}

	if input.Country != nil && len(*input.Country) > 0 {
		leagueToUpdate.Country = *input.Country
	}

	if input.Tier != nil && *input.Tier > 0 {
		leagueToUpdate.Tier = *input.Tier
	}

	leagueToUpdate = league.UpdateLeague(leagueToUpdate)

	return &model.League{
			ID:      leagueToUpdate.ID,
			Name:    leagueToUpdate.Name,
			Country: leagueToUpdate.Country,
			Tier:    leagueToUpdate.Tier},
		nil
}

func (r *mutationResolver) DeleteLeague(ctx context.Context, id int) (*model.DeleteResponse, error) {
	league.DeleteLeague(id)

	return &model.DeleteResponse{Message: "Successfully deleted league"}, nil
}

func (r *mutationResolver) CreateTeam(ctx context.Context, input model.NewTeam) (*model.Team, error) {
	var newTeam team.Team
	newTeam.League.ID = input.LeagueID
	newTeam.Name = input.Name
	newTeam.FoundingYear = input.FoundingYear

	newTeamID := newTeam.Save()

	newTeamLeague := league.GetLeagueById(input.LeagueID)

	return &model.Team{
			ID:           int(newTeamID),
			League:       (*model.League)(&newTeamLeague),
			Name:         newTeam.Name,
			FoundingYear: newTeam.FoundingYear},
		nil
}

func (r *mutationResolver) UpdateTeam(ctx context.Context, input model.UpdatedTeam) (*model.Team, error) {
	teamToUpdate := team.GetTeamById(input.ID)

	if input.LeagueID != nil && *input.LeagueID > 0 {
		teamToUpdate.League.ID = *input.LeagueID
	}

	if input.Name != nil && len(*input.Name) > 0 {
		teamToUpdate.Name = *input.Name
	}

	if input.FoundingYear != nil && *input.FoundingYear > 0 {
		teamToUpdate.FoundingYear = *input.FoundingYear
	}

	teamToUpdate = team.UpdateTeam(teamToUpdate)

	return &model.Team{
			ID:           teamToUpdate.ID,
			League:       (*model.League)(&teamToUpdate.League),
			Name:         teamToUpdate.Name,
			FoundingYear: teamToUpdate.FoundingYear},
		nil
}

func (r *mutationResolver) DeleteTeam(ctx context.Context, id int) (*model.DeleteResponse, error) {
	team.DeleteTeam(id)

	return &model.DeleteResponse{Message: "Successfully deleted team"}, nil
}

func (r *mutationResolver) CreatePlayer(ctx context.Context, input model.NewPlayer) (*model.Player, error) {
	var newPlayer player.Player
	newPlayer.FirstName = input.FirstName
	newPlayer.LastName = input.LastName
	newPlayer.Height = input.Height
	newPlayer.Nationality = input.Nationality
	newPlayer.Position = input.Position
	newPlayer.Team.ID = input.TeamID
	newPlayer.Number = input.Number
	newPlayer.Foot = input.Foot

	newPlayerID := newPlayer.Save()

	newPlayerTeam := team.GetTeamById(input.TeamID)

	return &model.Player{
			ID:          int(newPlayerID),
			FirstName:   newPlayer.FirstName,
			LastName:    newPlayer.LastName,
			Height:      newPlayer.Height,
			Nationality: newPlayer.Nationality,
			Position:    newPlayer.Position,
			Team: &model.Team{
				ID:           newPlayerTeam.ID,
				League:       (*model.League)(&newPlayerTeam.League),
				Name:         newPlayerTeam.Name,
				FoundingYear: newPlayerTeam.FoundingYear},
			Number: newPlayer.Number,
			Foot:   newPlayer.Foot},
		nil
}

func (r *mutationResolver) UpdatePlayer(ctx context.Context, input model.UpdatedPlayer) (*model.Player, error) {
	playerToUpdate := player.GetPlayerById(input.ID)

	if input.FirstName != nil && len(*input.FirstName) > 0 {
		playerToUpdate.FirstName = *input.FirstName
	}

	if input.LastName != nil && len(*input.LastName) > 0 {
		playerToUpdate.LastName = *input.LastName
	}

	if input.Height != nil && *input.Height > 0.00 {
		playerToUpdate.Height = *input.Height
	}

	if input.Nationality != nil && len(*input.Nationality) > 0 {
		playerToUpdate.Nationality = *input.Nationality
	}

	if input.Position != nil && len(*input.Position) > 0 {
		playerToUpdate.Position = *input.Position
	}

	if input.TeamID != nil && *input.TeamID > 0 {
		playerToUpdate.Team.ID = *input.TeamID
	}

	if input.Number != nil && *input.Number > 0 {
		playerToUpdate.Number = *input.Number
	}

	if input.Foot != nil && len(*input.Foot) > 0 {
		playerToUpdate.Foot = *input.Foot
	}

	playerToUpdate = player.UpdatePlayer(playerToUpdate)

	playerToUpdateTeam := team.GetTeamById(playerToUpdate.Team.ID)

	return &model.Player{
			ID:          playerToUpdate.ID,
			FirstName:   playerToUpdate.FirstName,
			LastName:    playerToUpdate.LastName,
			Height:      playerToUpdate.Height,
			Nationality: playerToUpdate.Nationality,
			Position:    playerToUpdate.Position,
			Team: &model.Team{ID: playerToUpdateTeam.ID,
				League:       (*model.League)(&playerToUpdateTeam.League),
				Name:         playerToUpdateTeam.Name,
				FoundingYear: playerToUpdateTeam.FoundingYear},
			Number: playerToUpdate.Number,
			Foot:   playerToUpdate.Foot},
		nil
}

func (r *mutationResolver) DeletePlayer(ctx context.Context, id int) (*model.DeleteResponse, error) {
	player.DeletePlayer(id)

	return &model.DeleteResponse{Message: "Successfully deleted player"}, nil
}

func (r *mutationResolver) CreateSeason(ctx context.Context, input model.NewSeason) (*model.Season, error) {
	var newSeason season.Season
	newSeason.League.ID = input.LeagueID
	newSeason.StartYear = input.StartYear
	newSeason.EndYear = input.EndYear

	newSeasonID := newSeason.Save()

	newSeason.League = league.GetLeagueById(input.LeagueID)

	return &model.Season{
			ID:        int(newSeasonID),
			League:    (*model.League)(&newSeason.League),
			StartYear: newSeason.StartYear,
			EndYear:   newSeason.EndYear},
		nil
}

func (r *mutationResolver) UpdateSeason(ctx context.Context, input model.UpdatedSeason) (*model.Season, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteSeason(ctx context.Context, id int) (*model.DeleteResponse, error) {
	season.DeleteSeason(id)

	return &model.DeleteResponse{Message: "Successfully deleted season"}, nil
}

func (r *mutationResolver) CreateMatch(ctx context.Context, input model.NewMatch) (*model.Match, error) {
	var newMatch match.Match
	newMatch.AwayTeam = team.GetTeamById(input.AwayTeamID)
	newMatch.HomeTeam = team.GetTeamById(input.HomeTeamID)
	newMatch.MatchDay = input.MatchDay
	newMatch.Season = season.GetSeasonById(input.SeasonID)

	newMatchID := newMatch.Save()

	return &model.Match{
			ID: int(newMatchID),
			Season: &model.Season{
				ID:        newMatch.Season.ID,
				League:    (*model.League)(&newMatch.Season.League),
				StartYear: newMatch.Season.StartYear,
				EndYear:   newMatch.Season.EndYear},
			HomeTeam: &model.Team{
				ID:           newMatch.HomeTeam.ID,
				League:       (*model.League)(&newMatch.HomeTeam.League),
				Name:         newMatch.HomeTeam.Name,
				FoundingYear: newMatch.HomeTeam.FoundingYear},
			AwayTeam: &model.Team{
				ID:           newMatch.AwayTeam.ID,
				League:       (*model.League)(&newMatch.AwayTeam.League),
				Name:         newMatch.AwayTeam.Name,
				FoundingYear: newMatch.AwayTeam.FoundingYear},
			MatchDay: newMatch.MatchDay},
		nil
}

func (r *mutationResolver) UpdateMatch(ctx context.Context, input model.UpdatedMatch) (*model.Match, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteMatch(ctx context.Context, id int) (*model.DeleteResponse, error) {
	match.DeleteMatch(id)

	return &model.DeleteResponse{Message: "Successfully deleted match"}, nil
}

func (r *mutationResolver) CreateMatchEvent(ctx context.Context, input model.NewMatchEvent) (*model.MatchEvent, error) {
	var newMatchEvent match.MatchEvent
	newMatchEvent.EventMinute = input.EventMinute
	newMatchEvent.EventType = match.MatchEventType(input.EventType)
	newMatchEvent.Match = match.GetMatchById(input.MatchID)
	newMatchEvent.Player = player.GetPlayerById(input.PlayerID)
	newMatchEvent.StoppageTime = *input.StoppageTime
	newMatchEvent.Team = team.GetTeamById(input.TeamID)

	newMatchEventID := newMatchEvent.Save()

	return &model.MatchEvent{ID: int(newMatchEventID),
			Match: &model.Match{
				ID: newMatchEvent.Match.ID,
				Season: &model.Season{
					ID:        newMatchEvent.Match.Season.ID,
					League:    (*model.League)(&newMatchEvent.Match.Season.League),
					StartYear: newMatchEvent.Match.Season.StartYear,
					EndYear:   newMatchEvent.Match.Season.EndYear},
				HomeTeam: &model.Team{
					ID:           newMatchEvent.Match.HomeTeam.ID,
					League:       (*model.League)(&newMatchEvent.Match.HomeTeam.League),
					Name:         newMatchEvent.Match.HomeTeam.Name,
					FoundingYear: newMatchEvent.Match.HomeTeam.FoundingYear},
				AwayTeam: &model.Team{
					ID:           newMatchEvent.Match.AwayTeam.ID,
					League:       (*model.League)(&newMatchEvent.Match.AwayTeam.League),
					Name:         newMatchEvent.Match.AwayTeam.Name,
					FoundingYear: newMatchEvent.Match.AwayTeam.FoundingYear},
				MatchDay: newMatchEvent.Match.MatchDay},
			Team: &model.Team{
				ID:           newMatchEvent.Team.ID,
				League:       (*model.League)(&newMatchEvent.Team.League),
				Name:         newMatchEvent.Team.Name,
				FoundingYear: newMatchEvent.Team.FoundingYear},
			Player: &model.Player{
				ID:          newMatchEvent.Player.ID,
				FirstName:   newMatchEvent.Player.FirstName,
				LastName:    newMatchEvent.Player.LastName,
				Height:      newMatchEvent.Player.Height,
				Nationality: newMatchEvent.Player.Nationality,
				Position:    newMatchEvent.Player.Position,
				Team: &model.Team{
					ID:           newMatchEvent.Player.Team.ID,
					League:       (*model.League)(&newMatchEvent.Player.Team.League),
					Name:         newMatchEvent.Player.Team.Name,
					FoundingYear: newMatchEvent.Player.Team.FoundingYear},
				Number: newMatchEvent.Player.Number, Foot: newMatchEvent.Player.Foot},
			EventType:    int(newMatchEvent.EventType),
			EventMinute:  newMatchEvent.EventMinute,
			StoppageTime: &newMatchEvent.StoppageTime},
		nil
}

func (r *mutationResolver) UpdateMatchEvent(ctx context.Context, input model.UpdatedMatchEvent) (*model.MatchEvent, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteMatchEvent(ctx context.Context, id int) (*model.DeleteResponse, error) {
	match.DeleteMatchEvent(id)

	return &model.DeleteResponse{Message: "Successfully deleted match event"}, nil
}

func (r *mutationResolver) CreateMatchStats(ctx context.Context, input model.NewMatchStats) (*model.MatchStats, error) {
	var newMatchStats match.MatchStats
	newMatchStats.CornersAway = input.CornersAway
	newMatchStats.CornersHome = input.CornersHome
	newMatchStats.FoulsAway = input.FoulsAway
	newMatchStats.FoulsHome = input.FoulsHome
	newMatchStats.MatchID = input.MatchID
	newMatchStats.OffsidesAway = input.OffsidesAway
	newMatchStats.OffsidesHome = input.OffsidesHome
	newMatchStats.PossessionAway = input.PossessionAway
	newMatchStats.PossessionHome = input.PossessionHome
	newMatchStats.SavesAway = input.SavesAway
	newMatchStats.SavesHome = input.SavesHome
	newMatchStats.ShotsOnTargetAway = input.ShotsOnTargetAway
	newMatchStats.ShotsOnTargetHome = input.ShotsOnTargetHome
	newMatchStats.TotalShotsAway = input.TotalShotsAway
	newMatchStats.TotalShotsHome = input.TotalShotsHome

	newMatchStats.ID = int(newMatchStats.Save())

	newMatchStatsMatch := match.GetMatchById(newMatchStats.MatchID)

	return &model.MatchStats{ID: newMatchStats.ID,
			Match: &model.Match{
				ID: newMatchStatsMatch.ID,
				Season: &model.Season{
					ID:        newMatchStatsMatch.Season.ID,
					League:    (*model.League)(&newMatchStatsMatch.Season.League),
					StartYear: newMatchStatsMatch.Season.StartYear,
					EndYear:   newMatchStatsMatch.Season.EndYear},
				HomeTeam: &model.Team{
					ID:           newMatchStatsMatch.HomeTeam.ID,
					League:       (*model.League)(&newMatchStatsMatch.HomeTeam.League),
					Name:         newMatchStatsMatch.HomeTeam.Name,
					FoundingYear: newMatchStatsMatch.HomeTeam.FoundingYear},
				AwayTeam: &model.Team{
					ID:           newMatchStatsMatch.AwayTeam.ID,
					League:       (*model.League)(&newMatchStatsMatch.AwayTeam.League),
					Name:         newMatchStatsMatch.AwayTeam.Name,
					FoundingYear: newMatchStatsMatch.AwayTeam.FoundingYear},
				MatchDay: newMatchStatsMatch.MatchDay},
			PossessionHome:    newMatchStats.PossessionHome,
			TotalShotsHome:    newMatchStats.TotalShotsHome,
			ShotsOnTargetHome: newMatchStats.ShotsOnTargetHome,
			SavesHome:         newMatchStats.SavesHome,
			CornersHome:       newMatchStats.CornersHome,
			FoulsHome:         newMatchStats.FoulsHome,
			OffsidesHome:      newMatchStats.OffsidesHome,
			PossessionAway:    newMatchStats.PossessionAway,
			TotalShotsAway:    newMatchStats.TotalShotsAway,
			ShotsOnTargetAway: newMatchStats.ShotsOnTargetAway,
			SavesAway:         newMatchStats.SavesAway,
			CornersAway:       newMatchStats.CornersAway,
			FoulsAway:         newMatchStats.FoulsAway,
			OffsidesAway:      newMatchStats.OffsidesAway},
		nil
}

func (r *mutationResolver) UpdateMatchStats(ctx context.Context, input model.UpdatedMatchStats) (*model.MatchStats, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteMatchStats(ctx context.Context, id int) (*model.DeleteResponse, error) {
	match.DeleteMatchStats(id)

	return &model.DeleteResponse{Message: "Successfully deleted match stats"}, nil
}

func (r *queryResolver) Leagues(ctx context.Context) ([]*model.League, error) {
	var resultLeagues []*model.League
	var dbLeagues []league.League
	dbLeagues = league.GetAllLeagues()

	for _, league := range dbLeagues {
		resultLeagues = append(resultLeagues,
			&model.League{
				ID:      league.ID,
				Name:    league.Name,
				Country: league.Country,
				Tier:    league.Tier})
	}

	return resultLeagues, nil
}

func (r *queryResolver) LeagueByID(ctx context.Context, leagueID int) (*model.League, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Teams(ctx context.Context) ([]*model.Team, error) {
	var resultTeams []*model.Team
	var dbTeams []team.Team
	dbTeams = team.GetAllTeams()

	for _, team := range dbTeams {

		resultTeams = append(resultTeams,
			&model.Team{
				ID:           team.ID,
				League:       (*model.League)(&team.League),
				Name:         team.Name,
				FoundingYear: team.FoundingYear})
	}

	return resultTeams, nil
}

func (r *queryResolver) TeamByID(ctx context.Context, teamID int) (*model.Team, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Players(ctx context.Context) ([]*model.Player, error) {
	var resultPlayers []*model.Player
	var dbPlayers []player.Player
	dbPlayers = player.GetAllPlayers()

	for _, player := range dbPlayers {

		// Need to figure out why casting the team field is not working like it is working for leagues
		resultPlayers = append(resultPlayers,
			&model.Player{
				ID:          player.ID,
				FirstName:   player.FirstName,
				LastName:    player.LastName,
				Height:      player.Height,
				Nationality: player.Nationality,
				Position:    player.Position,
				Team: &model.Team{
					ID:           player.Team.ID,
					League:       (*model.League)(&player.Team.League),
					Name:         player.Team.Name,
					FoundingYear: player.Team.FoundingYear},
				Number: player.Number,
				Foot:   player.Foot})
	}

	return resultPlayers, nil
}

func (r *queryResolver) PlayerByID(ctx context.Context, playerID int) (*model.Player, error) {
	playerResult := player.GetPlayerById(playerID)

	return &model.Player{
			ID:          playerResult.ID,
			FirstName:   playerResult.FirstName,
			LastName:    playerResult.LastName,
			Height:      playerResult.Height,
			Nationality: playerResult.Nationality,
			Position:    playerResult.Position,
			Team: &model.Team{
				ID:           playerResult.Team.ID,
				League:       (*model.League)(&playerResult.Team.League),
				Name:         playerResult.Team.Name,
				FoundingYear: playerResult.Team.FoundingYear},
			Number: playerResult.Number,
			Foot:   playerResult.Foot},
		nil
}

func (r *queryResolver) SeasonByLeagueID(ctx context.Context, leagueID int) ([]*model.Season, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SeasonByID(ctx context.Context, seasonID int) (*model.Season, error) {
	seasonResult := season.GetSeasonById(seasonID)

	return &model.Season{
			ID:        seasonResult.ID,
			League:    (*model.League)(&seasonResult.League),
			StartYear: seasonResult.StartYear,
			EndYear:   seasonResult.EndYear},
		nil
}

func (r *queryResolver) Matches(ctx context.Context) ([]*model.Match, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MatchByID(ctx context.Context, matchID int) (*model.Match, error) {
	matchResult := match.GetMatchById(matchID)

	return &model.Match{
			ID: matchResult.ID,
			Season: &model.Season{
				ID:        matchResult.Season.ID,
				League:    (*model.League)(&matchResult.Season.League),
				StartYear: matchResult.Season.StartYear,
				EndYear:   matchResult.Season.EndYear},
			HomeTeam: &model.Team{
				ID:           matchResult.HomeTeam.ID,
				League:       (*model.League)(&matchResult.HomeTeam.League),
				Name:         matchResult.HomeTeam.Name,
				FoundingYear: matchResult.HomeTeam.FoundingYear},
			AwayTeam: &model.Team{
				ID:           matchResult.AwayTeam.ID,
				League:       (*model.League)(&matchResult.AwayTeam.League),
				Name:         matchResult.AwayTeam.Name,
				FoundingYear: matchResult.AwayTeam.FoundingYear},
			MatchDay: matchResult.MatchDay},
		nil
}

func (r *queryResolver) MatchesBySeasonID(ctx context.Context, seasonID int) ([]*model.Match, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MatchEventsByMatchID(ctx context.Context, matchID int) ([]*model.MatchEvent, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MatchStatsByMatchID(ctx context.Context, matchID int) (*model.MatchStats, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
