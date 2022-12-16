package test

import (
	"MLBTakeHome/schema"
	"MLBTakeHome/sortgames"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var defaultTeamId = 134
var notOurTeam1 = 130
var notOurTeam2 = 135
var notOurTeam3 = 150
var notOurTeam4 = 151
var notOurTeam5 = 152

func TestSortGameDifferingTeam(t *testing.T) {
	type args struct {
		respSchedule *schema.Schedule
		teamId       int
	}
	tests := []struct {
		name string
		args args
		want []schema.Game
	}{
		{
			name: "NoGames",
			args: args{
				respSchedule: &schema.Schedule{
					Dates: []schema.Date{
						{
							Games: []schema.Game{},
						},
					},
				},
				teamId: defaultTeamId,
			},
			want: []schema.Game{},
		},
		{
			name: "OneGameMatches",
			args: args{
				respSchedule: &schema.Schedule{
					Dates: []schema.Date{
						{
							Games: []schema.Game{
								CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam2).GetGame(),
							},
						},
					},
				},
				teamId: defaultTeamId,
			},
			want: []schema.Game{
				CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam2).GetGame(),
			},
		},
		{
			name: "OneGameDoesntMatch",
			args: args{
				respSchedule: &schema.Schedule{
					Dates: []schema.Date{
						{
							Games: []schema.Game{
								CreateGame().HomeTeam(notOurTeam1).AwayTeam(notOurTeam2).GetGame(),
							},
						},
					},
				},
				teamId: defaultTeamId,
			},
			want: []schema.Game{
				CreateGame().HomeTeam(notOurTeam1).AwayTeam(notOurTeam2).GetGame(),
			},
		},
		{
			name: "TwoGamesOneMatchesAlreadyInOrder",
			args: args{
				respSchedule: &schema.Schedule{
					Dates: []schema.Date{
						{
							Games: []schema.Game{
								CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GetGame(),
								CreateGame().HomeTeam(notOurTeam2).AwayTeam(notOurTeam3).GetGame(),
							},
						},
					},
				},
				teamId: defaultTeamId,
			},
			want: []schema.Game{
				CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GetGame(),
				CreateGame().HomeTeam(notOurTeam2).AwayTeam(notOurTeam3).GetGame(),
			},
		},
		{
			name: "TwoGamesOneMatchesOutOfOrder",
			args: args{
				respSchedule: &schema.Schedule{
					Dates: []schema.Date{
						{
							Games: []schema.Game{
								CreateGame().HomeTeam(notOurTeam2).AwayTeam(notOurTeam3).GetGame(),
								CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GetGame(),
							},
						},
					},
				},
				teamId: defaultTeamId,
			},
			want: []schema.Game{
				CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GetGame(),
				CreateGame().HomeTeam(notOurTeam2).AwayTeam(notOurTeam3).GetGame(),
			},
		},
		{
			name: "ThreeGamesOutOfOrder",
			args: args{
				respSchedule: &schema.Schedule{
					Dates: []schema.Date{
						{
							Games: []schema.Game{
								CreateGame().HomeTeam(notOurTeam2).AwayTeam(notOurTeam3).GetGame(),
								CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GetGame(),
								CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam2).GetGame(),
							},
						},
					},
				},
				teamId: defaultTeamId,
			},
			want: []schema.Game{
				CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GetGame(),
				CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam2).GetGame(),
				CreateGame().HomeTeam(notOurTeam2).AwayTeam(notOurTeam3).GetGame(),
			},
		},
		{
			name: "MultipleGamesNoneMatch",
			args: args{
				respSchedule: &schema.Schedule{
					Dates: []schema.Date{
						{
							Games: []schema.Game{
								CreateGame().HomeTeam(notOurTeam1).AwayTeam(notOurTeam2).GetGame(),
								CreateGame().HomeTeam(notOurTeam2).AwayTeam(notOurTeam3).GetGame(),
								CreateGame().HomeTeam(notOurTeam3).AwayTeam(notOurTeam1).GetGame(),
							},
						},
					},
				},
				teamId: defaultTeamId,
			},
			want: []schema.Game{
				CreateGame().HomeTeam(notOurTeam1).AwayTeam(notOurTeam2).GetGame(),
				CreateGame().HomeTeam(notOurTeam2).AwayTeam(notOurTeam3).GetGame(),
				CreateGame().HomeTeam(notOurTeam3).AwayTeam(notOurTeam1).GetGame(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortgames.SortGames(tt.args.respSchedule, tt.args.teamId)
			got := tt.args.respSchedule.Dates[0].Games

			gotBytes, _ := json.Marshal(got)
			wantBytes, _ := json.Marshal(tt.want)

			assert.Equal(t, string(wantBytes), string(gotBytes), "not equal")
		})
	}
}

func TestSortGameDifferingTime(t *testing.T) {
	past := time.Now().Add(-1 * time.Hour)
	now := time.Now()
	future := time.Now().Add(5 * time.Hour)
	type args struct {
		respSchedule *schema.Schedule
		teamId       int
	}
	tests := []struct {
		name string
		args args
		want []schema.Game
	}{
		{
			name: "TwoGamesSameTeamDifferentTime",
			args: args{
				respSchedule: &schema.Schedule{
					Dates: []schema.Date{
						{
							Games: []schema.Game{
								CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GameDate(now).GetGame(),
								CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GameDate(past).GetGame(),
							},
						},
					},
				},
				teamId: defaultTeamId,
			},
			want: []schema.Game{
				CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GameDate(past).GetGame(),
				CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GameDate(now).GetGame(),
			},
		},
		{
			name: "ThreeGamesSameTeamDifferentTimes",
			args: args{
				respSchedule: &schema.Schedule{
					Dates: []schema.Date{
						{
							Games: []schema.Game{
								CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GameDate(now).GetGame(),
								CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GameDate(future).GetGame(),
								CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GameDate(past).GetGame(),
							},
						},
					},
				},
				teamId: defaultTeamId,
			},
			want: []schema.Game{
				CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GameDate(past).GetGame(),
				CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GameDate(now).GetGame(),
				CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GameDate(future).GetGame(),
			},
		},
		//{
		//	name: "ThreeGamesDifferentTimesAndTeams",
		//	args: args{
		//		respSchedule: &schema.Schedule{
		//			Dates: []schema.Date{
		//				{
		//					Games: []schema.Game{
		//						makeGame(notOurTeam2, notOurTeam1, past, false, false, false, 3),
		//						makeGame(notOurTeam3, defaultTeamId, now, false, false, false, 1),
		//						makeGame(notOurTeam3, defaultTeamId, future, false, false, false, 2),
		//					},
		//				},
		//			},
		//		},
		//		teamId: defaultTeamId,
		//	},
		//	want: []schema.Game{
		//		makeGame(notOurTeam3, defaultTeamId, now, false, false, false, 1),
		//		makeGame(notOurTeam3, defaultTeamId, future, false, false, false, 2),
		//		makeGame(notOurTeam2, notOurTeam1, past, false, false, false, 3),
		//	},
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortgames.SortGames(tt.args.respSchedule, tt.args.teamId)
			got := tt.args.respSchedule.Dates[0].Games

			gotBytes, _ := json.Marshal(got)
			wantBytes, _ := json.Marshal(tt.want)

			assert.Equal(t, string(wantBytes), string(gotBytes), "not equal")
		})
	}
}

func TestSortGameDoubleheader(t *testing.T) {
	now := time.Now()
	type args struct {
		respSchedule *schema.Schedule
		teamId       int
	}
	tests := []struct {
		name string
		args args
		want []schema.Game
	}{
		{
			name: "SameStartTimeOneTBD",
			args: args{
				respSchedule: &schema.Schedule{
					Dates: []schema.Date{
						{
							Games: []schema.Game{
								CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GameDate(now).IsRegDoubleheader().TimeTBD().GetGame(),
								CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GameDate(now).IsRegDoubleheader().GetGame(),
							},
						},
					},
				},
				teamId: defaultTeamId,
			},
			want: []schema.Game{
				CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GameDate(now).IsRegDoubleheader().GetGame(),
				CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GameDate(now).IsRegDoubleheader().TimeTBD().GetGame(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortgames.SortGames(tt.args.respSchedule, tt.args.teamId)
			got := tt.args.respSchedule.Dates[0].Games

			gotBytes, _ := json.Marshal(got)
			wantBytes, _ := json.Marshal(tt.want)

			assert.Equal(t, string(wantBytes), string(gotBytes), "game array not properly sorted")
		})
	}
}

func TestSortGameLiveGame(t *testing.T) {
	past := time.Now().Add(-1 * time.Hour)
	now := time.Now()
	future := time.Now().Add(5 * time.Hour)
	type args struct {
		respSchedule *schema.Schedule
		teamId       int
	}
	tests := []struct {
		name string
		args args
		want []schema.Game
	}{
		{
			name: "LiveWithDoubleHeaders",
			args: args{
				respSchedule: &schema.Schedule{
					Dates: []schema.Date{
						{
							Games: []schema.Game{

								CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GameDate(now).
									IsRegDoubleheader().GetGame(),
								CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GameDate(now).
									IsLive().GetGame(),
								CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GameDate(now).
									IsRegDoubleheader().TimeTBD().GetGame(),
							},
						},
					},
				},
				teamId: defaultTeamId,
			},
			want: []schema.Game{
				CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GameDate(now).
					IsLive().GetGame(),
				CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GameDate(now).
					IsRegDoubleheader().GetGame(),
				CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GameDate(now).
					IsRegDoubleheader().TimeTBD().GetGame(),
			},
		},

		{
			name: "OneGameLive",
			args: args{
				respSchedule: &schema.Schedule{
					Dates: []schema.Date{
						{
							Games: []schema.Game{
								CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GameDate(past).
									GetGame(),
								CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GameDate(now).
									IsLive().GetGame(),
								CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GameDate(future).
									GetGame(),
							},
						},
					},
				},
				teamId: defaultTeamId,
			},
			want: []schema.Game{
				CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GameDate(now).
					IsLive().GetGame(),
				CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GameDate(past).
					GetGame(),
				CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GameDate(future).
					GetGame(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortgames.SortGames(tt.args.respSchedule, tt.args.teamId)
			got := tt.args.respSchedule.Dates[0].Games

			gotBytes, _ := json.Marshal(got)
			wantBytes, _ := json.Marshal(tt.want)

			assert.Equal(t, string(wantBytes), string(gotBytes), "game array not properly sorted")
		})
	}
}

func TestSortGameMultipleCriteria(t *testing.T) {
	past := time.Now().Add(-1 * time.Hour)
	now := time.Now()
	future := time.Now().Add(5 * time.Hour)
	type args struct {
		respSchedule *schema.Schedule
		teamId       int
	}
	tests := []struct {
		name string
		args args
		want []schema.Game
	}{
		{
			name: "DifferentTimesAndTeams",
			args: args{
				respSchedule: &schema.Schedule{
					Dates: []schema.Date{
						{
							Games: []schema.Game{
								CreateGame().HomeTeam(notOurTeam2).AwayTeam(notOurTeam3).GameDate(now).GetGame(),
								CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GameDate(future).GetGame(),
								CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GameDate(now).GetGame(),
							},
						},
					},
				},
				teamId: defaultTeamId,
			},
			want: []schema.Game{
				CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GameDate(now).GetGame(),
				CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GameDate(future).GetGame(),
				CreateGame().HomeTeam(notOurTeam2).AwayTeam(notOurTeam3).GameDate(now).GetGame(),
			},
		},

		{
			name: "TeamsTimesDoubleheaderLive",
			args: args{
				respSchedule: &schema.Schedule{
					Dates: []schema.Date{
						{
							Games: []schema.Game{
								CreateGame().HomeTeam(notOurTeam2).AwayTeam(notOurTeam3).
									GameDate(past).IsRegDoubleheader().TimeTBD().IsLive().GetGame(),
								CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GameDate(now).
									IsRegDoubleheader().TimeTBD().GetGame(),
								CreateGame().HomeTeam(notOurTeam2).AwayTeam(notOurTeam3).
									GameDate(past).IsRegDoubleheader().IsLive().GetGame(),
								CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GameDate(now).
									IsRegDoubleheader().GetGame(),
								CreateGame().HomeTeam(notOurTeam4).AwayTeam(notOurTeam5).GameDate(past).GetGame(),
								CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GameDate(future).IsLive().GetGame(),
								CreateGame().HomeTeam(notOurTeam5).AwayTeam(notOurTeam4).GameDate(now).GetGame(),
							},
						},
					},
				},
				teamId: defaultTeamId,
			},
			want: []schema.Game{

				CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GameDate(future).IsLive().GetGame(),
				CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GameDate(now).
					IsRegDoubleheader().GetGame(),
				CreateGame().HomeTeam(defaultTeamId).AwayTeam(notOurTeam1).GameDate(now).
					IsRegDoubleheader().TimeTBD().GetGame(),

				CreateGame().HomeTeam(notOurTeam2).AwayTeam(notOurTeam3).
					GameDate(past).IsRegDoubleheader().TimeTBD().IsLive().GetGame(),
				CreateGame().HomeTeam(notOurTeam2).AwayTeam(notOurTeam3).
					GameDate(past).IsRegDoubleheader().IsLive().GetGame(),
				CreateGame().HomeTeam(notOurTeam4).AwayTeam(notOurTeam5).GameDate(past).GetGame(),
				CreateGame().HomeTeam(notOurTeam5).AwayTeam(notOurTeam4).GameDate(now).GetGame(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortgames.SortGames(tt.args.respSchedule, tt.args.teamId)
			got := tt.args.respSchedule.Dates[0].Games

			gotBytes, _ := json.Marshal(got)
			wantBytes, _ := json.Marshal(tt.want)

			assert.Equal(t, string(wantBytes), string(gotBytes), "game array not properly sorted")
		})
	}
}
