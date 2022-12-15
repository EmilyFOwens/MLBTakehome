package sortgames

import (
	"MLBTakeHome/schema"
)

func SortGameFunction(respSchedule *schema.Schedule, teamId int) {
	gameSorter := GameSorter{
		games:  respSchedule.Dates[0].Games,
		teamId: teamId,
	}

	gameSorter.Sort()
}
