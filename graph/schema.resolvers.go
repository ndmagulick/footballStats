package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"footballStats/graph/generated"
	"footballStats/graph/model"
	league "footballStats/leagues"
	player "footballStats/players"
	team "footballStats/teams"
)

func (r *mutationResolver) CreateLeague(ctx context.Context, input model.NewLeague) (*model.League, error) {
	var newLeague league.League
	newLeague.Name = input.Name
	newLeague.Country = input.Country
	newLeague.Tier = input.Tier

	newLeagueID := newLeague.Save()

	return &model.League{ID: int(newLeagueID), Name: newLeague.Name, Country: newLeague.Country, Tier: newLeague.Tier}, nil
}

func (r *mutationResolver) CreateTeam(ctx context.Context, input model.NewTeam) (*model.Team, error) {
	var newTeam team.Team
	newTeam.League.ID = input.LeagueID
	newTeam.Name = input.Name
	newTeam.FoundingYear = input.FoundingYear

	newTeamID := newTeam.Save()

	newTeamLeague := league.GetLeagueById(input.LeagueID)

	return &model.Team{ID: int(newTeamID), League: (*model.League)(&newTeamLeague), Name: newTeam.Name, FoundingYear: newTeam.FoundingYear}, nil
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

	return &model.Player{ID: int(newPlayerID), FirstName: newPlayer.FirstName, LastName: newPlayer.LastName, Height: newPlayer.Height, Nationality: newPlayer.Nationality,
		Position: newPlayer.Position, Team: &model.Team{ID: newPlayerTeam.ID, League: (*model.League)(&newPlayerTeam.League), Name: newPlayerTeam.Name,
			FoundingYear: newPlayerTeam.FoundingYear}, Number: newPlayer.Number, Foot: newPlayer.Foot}, nil
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

	return &model.League{ID: leagueToUpdate.ID, Name: leagueToUpdate.Name, Country: leagueToUpdate.Country, Tier: leagueToUpdate.Tier}, nil
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

	return &model.Team{ID: teamToUpdate.ID, League: (*model.League)(&teamToUpdate.League), Name: teamToUpdate.Name, FoundingYear: teamToUpdate.FoundingYear}, nil
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

	return &model.Player{ID: playerToUpdate.ID, FirstName: playerToUpdate.FirstName, LastName: playerToUpdate.LastName, Height: playerToUpdate.Height,
		Nationality: playerToUpdate.Nationality, Position: playerToUpdate.Position, Team: &model.Team{ID: playerToUpdateTeam.ID,
			League: (*model.League)(&playerToUpdateTeam.League), Name: playerToUpdateTeam.Name, FoundingYear: playerToUpdateTeam.FoundingYear},
		Number: playerToUpdate.Number, Foot: playerToUpdate.Foot}, nil
}

func (r *mutationResolver) DeleteLeague(ctx context.Context, id int) (*model.DeleteResponse, error) {
	league.DeleteLeague(id)

	return &model.DeleteResponse{Message: "Successfully deleted league"}, nil
}

func (r *mutationResolver) DeleteTeam(ctx context.Context, id int) (*model.DeleteResponse, error) {
	team.DeleteTeam(id)

	return &model.DeleteResponse{Message: "Successfully deleted team"}, nil
}

func (r *mutationResolver) DeletePlayer(ctx context.Context, id int) (*model.DeleteResponse, error) {
	player.DeletePlayer(id)

	return &model.DeleteResponse{Message: "Successfully deleted player"}, nil
}

func (r *queryResolver) Leagues(ctx context.Context) ([]*model.League, error) {
	var resultLeagues []*model.League
	var dbLeagues []league.League
	dbLeagues = league.GetAllLeagues()

	for _, league := range dbLeagues {
		resultLeagues = append(resultLeagues, &model.League{ID: league.ID, Name: league.Name, Country: league.Country, Tier: league.Tier})
	}

	return resultLeagues, nil
}

func (r *queryResolver) Teams(ctx context.Context) ([]*model.Team, error) {
	var resultTeams []*model.Team
	var dbTeams []team.Team
	dbTeams = team.GetAllTeams()

	for _, team := range dbTeams {

		resultTeams = append(resultTeams, &model.Team{ID: team.ID, League: (*model.League)(&team.League), Name: team.Name, FoundingYear: team.FoundingYear})
	}

	return resultTeams, nil
}

func (r *queryResolver) Players(ctx context.Context) ([]*model.Player, error) {
	var resultPlayers []*model.Player
	var dbPlayers []player.Player
	dbPlayers = player.GetAllPlayers()

	for _, player := range dbPlayers {

		// Need to figure out why casting the team field is not working like it is working for leagues
		resultPlayers = append(resultPlayers, &model.Player{ID: player.ID, FirstName: player.FirstName, LastName: player.LastName, Height: player.Height,
			Nationality: player.Nationality, Position: player.Position, Team: &model.Team{ID: player.Team.ID, League: (*model.League)(&player.Team.League), Name: player.Team.Name,
				FoundingYear: player.Team.FoundingYear}, Number: player.Number, Foot: player.Foot})
	}

	return resultPlayers, nil
}

func (r *queryResolver) Seasons(ctx context.Context) ([]*model.Season, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Matches(ctx context.Context) ([]*model.Match, error) {
	panic(fmt.Errorf("not implemented"))
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
