package test

import (
	"MLBTakeHome/schema"
	"MLBTakeHome/sortgames"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestSortedGames(t *testing.T) {
	type args struct {
		teamId int
		date   string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Basic",
			args: args{
				teamId: 133,
				date:   "202d",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var respScheduleOriginal schema.Schedule
			var respScheduleSorted schema.Schedule
			got, _ := sortgames.FavoriteTeamGames(tt.args.teamId, tt.args.date)

			_ = json.Unmarshal(got, &respScheduleSorted)

			for _, x := range respScheduleSorted.Dates {
				for _, y := range x.Games {
					fmt.Println(y)
				}
			}

			// Get games by date
			resp, err := http.Get(fmt.Sprintf("https://statsapi.mlb.com/api/v1/schedule?date=%s&sportId=1&language=en", tt.args.date))
			defer resp.Body.Close()
			if err != nil {
				return
			}

			err = json.NewDecoder(resp.Body).Decode(&respScheduleOriginal)
			if err != nil {
				return
			}

			if len(respScheduleOriginal.Dates) != 0 {
				assert.ElementsMatch(t, respScheduleOriginal.Dates[0].Games, respScheduleSorted.Dates[0].Games)
			}
		})
	}
}
