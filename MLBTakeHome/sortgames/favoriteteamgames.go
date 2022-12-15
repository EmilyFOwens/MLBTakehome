package sortgames

import (
	"MLBTakeHome/schema"
	"encoding/json"
	"fmt"
	"net/http"
)

// FavoriteTeamGames sends the list of newly sorted games to the client
//
// As this is being treated as proof of concept, errors aren't handled beyond being passed to the
// calling function; For the purpose of this exercise, the assumption is that the client will
// itself have some sort of error handling
//
// Returns json as []byte
func FavoriteTeamGames(teamId int, date string) ([]byte, error) {
	var respSchedule schema.Schedule

	// Get games by date
	resp, err := http.Get(fmt.Sprintf("https://statsapi.mlb.com/api/v1/schedule?date=%s&sportId=1&language=en", date))
	defer resp.Body.Close()
	if err != nil {
		return []byte{}, err
	}

	err = json.NewDecoder(resp.Body).Decode(&respSchedule)
	if err != nil {
		return []byte{}, err
	}

	// Sort retrieved schedule
	SortGameFunction(&respSchedule, teamId)

	// Return sorted schedule
	return json.Marshal(respSchedule)
}
