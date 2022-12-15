package test

import (
	"MLBTakeHome/schema"
	"MLBTakeHome/sortgames"
	"time"
)

// makeGame generates a schema.Game for testing purposes
// There are a limited number of fields relevant for our sorting purposes, so this
// function only builds games with these values present
// Note: Given very limited time, and the knowledge that this is to be considered a prototype, this
// is functional but perhaps not the ideal way of building this test struct. Additional care would be taken in a
// production/real testing environment
func makeGame(homeId, awayId int, gameDate time.Time, doubleheader, tbd, live bool, gamePk int) schema.Game {
	var gameStatus string
	var doubleheaderStatus string
	if live {
		gameStatus = sortgames.LiveGameStatusCode
	}
	if doubleheader {
		doubleheaderStatus = sortgames.RegularDoubleheader
	}
	return schema.Game{
		GamePk:   gamePk,
		GameDate: gameDate,
		Status: schema.Status{
			StatusCode:   gameStatus,
			StartTimeTBD: tbd,
		},
		DoubleHeader: doubleheaderStatus,
		Teams: schema.Teams{
			Away: schema.Away{
				Team: schema.Team{
					Id: awayId,
				},
			},
			Home: schema.Home{
				Team: schema.Team{
					Id: homeId,
				},
			},
		},
	}
}
