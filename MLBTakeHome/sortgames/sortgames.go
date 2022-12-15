package sortgames

import (
	"MLBTakeHome/schema"
)

func SortGames(respSchedule *schema.Schedule, teamId int) {
	if respSchedule.Dates == nil {
		return
	}
	gameSorter := GameSorter{
		games:  respSchedule.Dates[0].Games,
		teamId: teamId,
	}

	gameSorter.Sort()
}
