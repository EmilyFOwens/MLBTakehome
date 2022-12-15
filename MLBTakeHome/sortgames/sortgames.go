package sortgames

import (
	"MLBTakeHome/schema"
)

// SortGames creates the GameSorter for the given team and schedule, then calls Sort()
func SortGames(respSchedule *schema.Schedule, teamId int) {
	if len(respSchedule.Dates) == 0 {
		return
	}

	gameSorter := GameSorter{
		games:  respSchedule.Dates[0].Games,
		teamId: teamId,
	}

	gameSorter.Sort()
}
