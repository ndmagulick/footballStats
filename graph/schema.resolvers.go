package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
