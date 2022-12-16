package test

import (
	"MLBTakeHome/schema"
	"MLBTakeHome/sortgames"
	"time"
)

// GameBuilder is the base for creating test game data
type GameBuilder struct {
	Game schema.Game
}

// CreateGame initializes new GameBuilder
func CreateGame() *GameBuilder {
	return &GameBuilder{Game: schema.Game{}}
}

// HomeTeam sets the game's home team to given id
func (gb *GameBuilder) HomeTeam(teamId int) *GameBuilder {
	gb.Game.Teams.Home.Team.Id = teamId
	return gb
}

// AwayTeam sets the game's away team to given id
func (gb *GameBuilder) AwayTeam(teamId int) *GameBuilder {
	gb.Game.Teams.Away.Team.Id = teamId
	return gb
}

// GameDate sets the game's game date to given time
func (gb *GameBuilder) GameDate(gameDate time.Time) *GameBuilder {
	gb.Game.GameDate = gameDate
	return gb
}

// IsRegDoubleheader sets the game's doubleheader state to sortgames.RegularDoubleheader
func (gb *GameBuilder) IsRegDoubleheader() *GameBuilder {
	gb.Game.DoubleHeader = sortgames.RegularDoubleheader
	return gb
}

// IsLive sets the game's status code to sortgames.LiveGameStatusCode
// TODO: This may not be the right value, need more documentation to be sure
func (gb *GameBuilder) IsLive() *GameBuilder {
	gb.Game.Status.StatusCode = sortgames.LiveGameStatusCode
	return gb
}

// TimeTBD sets the game's start time tbd value to true
func (gb *GameBuilder) TimeTBD() *GameBuilder {
	gb.Game.Status.StartTimeTBD = true
	return gb
}

// GetGame returns the fully built game
func (gb *GameBuilder) GetGame() schema.Game {
	return gb.Game
}
