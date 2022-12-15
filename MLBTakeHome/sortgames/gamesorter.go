package sortgames

import (
	"MLBTakeHome/schema"
	"sort"
)

type comparisonFunc func(game1, game2 *schema.Game) bool

// GameSorter sorts games, favoring games with the teamId of
// the given favorite team
type GameSorter struct {
	games           []schema.Game
	comparisonFuncs []comparisonFunc
	teamId          int
}

// OrderedBy appends any new comparison functions to the GameSorter's
// list of comparison functions
// Note: gs.Sort() has a set of default comparators, which are built inside gs.Sort(), using this method
func (gs *GameSorter) OrderedBy(comp ...comparisonFunc) {
	gs.comparisonFuncs = append(gs.comparisonFuncs, comp...)
}

// Len implements the sort interface Len() method
func (gs *GameSorter) Len() int {
	return len(gs.games)
}

// Swap implements the sort interface Swap() method
func (gs *GameSorter) Swap(i, j int) {
	gs.games[i], gs.games[j] = gs.games[j], gs.games[i]
}

// Less implements the sort interface Less() method, and is used to sort schema.Game objects
// If game1 is 'less' than game2, then game1 will be listed before the game 2, and vice versa
func (gs *GameSorter) Less(i, j int) bool {
	game1, game2 := &gs.games[i], &gs.games[j]
	var k int
	for k = 0; k < len(gs.comparisonFuncs)-1; k++ {
		less := gs.comparisonFuncs[k]
		switch {
		case less(game1, game2):
			return true
		case less(game2, game1):
			return false
		}
	}

	return gs.comparisonFuncs[k](game1, game2)
}

// Sort implements the sort.Sort() method
// This method adds the basic comparators needed to properly sort
// the games as defined in the problem definition
// Stable sort here to guarantee preservation of elements which don't include
// favorite team
// This is less efficient for larger datasets, but the maximum number of games
// in a day is very small
func (gs *GameSorter) Sort() {
	gs.OrderedBy(gs.byFavoriteTeam, gs.byLiveGame, gs.byGameStartTime)
	sort.Stable(gs)
}

// byFavoriteTeam sorts the games based on the defined "favorite team"
// Game1 is required to be listed before (aka is less than) Game2 only under the following conditions:
// Game 1 must include the favorite team
// Game 2 must not include the favorite team
func (gs *GameSorter) byFavoriteTeam(game1, game2 *schema.Game) bool {
	if game1 == nil || game2 == nil {
		return false
	}
	if favoriteTeamIsPlaying(game1, gs.teamId) {
		if !favoriteTeamIsPlaying(game2, gs.teamId) {
			return true
		}
	}
	return false
}

// byGameStartTime sorts the games considering both the gameDate value,
// and if the games in consideration are part of a doubleheader
func (gs *GameSorter) byGameStartTime(game1, game2 *schema.Game) bool {
	if game1 == nil || game2 == nil {
		return false
	}

	// If it isn't a game by our favorite team, there is no need to perform any more operations
	if !favoriteTeamIsPlaying(game1, gs.teamId) && !favoriteTeamIsPlaying(game2, gs.teamId) {
		return false
	}

	// If normal doubleheader, both games should have DoubleHeader set to "Y"
	// In this case, the game with StartTimeTBD equal to true is the later game
	if game1.DoubleHeader == RegularDoubleheader && game2.DoubleHeader == RegularDoubleheader {
		if game2.Status.StartTimeTBD == true {
			return true
		}
	}

	// If single game, or split doubleheader, we can just compare GameDate
	if game1.GameDate.Before(game2.GameDate) {
		return true
	}

	return false
}

// byLiveGame sorts games such that a live game is displayed before any non-live game
func (gs *GameSorter) byLiveGame(game1, game2 *schema.Game) bool {
	if game1 == nil || game2 == nil {
		return false
	}

	if !favoriteTeamIsPlaying(game1, gs.teamId) || !favoriteTeamIsPlaying(game2, gs.teamId) {
		return false
	}

	// There are no examples of a Live game to be found in the dataset,
	// so this is a guess at an identifier based on observation of other
	// game statuses
	if game1.Status.StatusCode == LiveGameStatusCode && game2.Status.StatusCode != LiveGameStatusCode {
		return true
	}

	return false
}

// favoriteTeamIsPlaying checks to see if either the home or away team is
// our favorite team
func favoriteTeamIsPlaying(game *schema.Game, teamId int) bool {
	if game == nil {
		return false
	}
	if game.Teams.Home.Team.Id == teamId || game.Teams.Away.Team.Id == teamId {
		return true
	}
	return false
}
