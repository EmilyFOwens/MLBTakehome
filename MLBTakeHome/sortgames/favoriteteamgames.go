package sortgames

import (
	"MLBTakeHome/schema"
	"encoding/json"
	"fmt"
	"net/http"
)

// FavoriteTeamGames sends the list of newly sorted games to the client
// If no http error, sorts games and returns json as []byte
// If http error occurs, returns http status code as []byte
func FavoriteTeamGames(teamId int, date string) ([]byte, error) {
	var respSchedule schema.Schedule

	// Get games by date
	resp, err := http.Get(fmt.Sprintf("https://statsapi.mlb.com/api/v1/schedule?date=%s&sportId=1&language=en", date))
	defer resp.Body.Close()
	if err != nil {
		return []byte{}, err
	}

	if resp.StatusCode >= 300 {
		return []byte(resp.Status), nil
	}

	err = json.NewDecoder(resp.Body).Decode(&respSchedule)
	if err != nil {
		return []byte{}, err
	}

	// Sort retrieved schedule
	SortGames(&respSchedule, teamId)

	// Return sorted schedule
	return json.Marshal(respSchedule)
}
