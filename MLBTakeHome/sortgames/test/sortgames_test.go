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

func TestSortGameDifferingTeam(t *testing.T) {
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
								makeGame(notOurTeam1, defaultTeamId, now, false, false, false, 1),
							},
						},
					},
				},
				teamId: defaultTeamId,
			},
			want: []schema.Game{
				makeGame(notOurTeam1, defaultTeamId, now, false, false, false, 1),
			},
		},
		{
			name: "OneGameDoesntMatch",
			args: args{
				respSchedule: &schema.Schedule{
					Dates: []schema.Date{
						{
							Games: []schema.Game{
								makeGame(notOurTeam2, notOurTeam1, now, false, false, false, 1),
							},
						},
					},
				},
				teamId: defaultTeamId,
			},
			want: []schema.Game{
				makeGame(notOurTeam2, notOurTeam1, now, false, false, false, 1),
			},
		},
		{
			name: "TwoGamesOneMatchesAlreadyInOrder",
			args: args{
				respSchedule: &schema.Schedule{
					Dates: []schema.Date{
						{
							Games: []schema.Game{
								makeGame(notOurTeam1, defaultTeamId, now, false, false, false, 1),
								makeGame(notOurTeam3, notOurTeam2, now, false, false, false, 2),
							},
						},
					},
				},
				teamId: defaultTeamId,
			},
			want: []schema.Game{
				makeGame(notOurTeam1, defaultTeamId, now, false, false, false, 1),
				makeGame(notOurTeam3, notOurTeam2, now, false, false, false, 2),
			},
		},
		{
			name: "TwoGamesOneMatchesOutOfOrder",
			args: args{
				respSchedule: &schema.Schedule{
					Dates: []schema.Date{
						{
							Games: []schema.Game{
								makeGame(notOurTeam3, notOurTeam2, now, false, false, false, 2),
								makeGame(notOurTeam1, defaultTeamId, now, false, false, false, 1),
							},
						},
					},
				},
				teamId: defaultTeamId,
			},
			want: []schema.Game{
				makeGame(notOurTeam1, defaultTeamId, now, false, false, false, 1),
				makeGame(notOurTeam3, notOurTeam2, now, false, false, false, 2),
			},
		},
		{
			name: "ThreeGamesOutOfOrder",
			args: args{
				respSchedule: &schema.Schedule{
					Dates: []schema.Date{
						{
							Games: []schema.Game{
								makeGame(notOurTeam2, notOurTeam1, now, false, false, false, 3),
								makeGame(notOurTeam3, defaultTeamId, now, false, false, false, 1),
								makeGame(notOurTeam2, defaultTeamId, now, false, false, false, 2),
							},
						},
					},
				},
				teamId: defaultTeamId,
			},
			want: []schema.Game{
				makeGame(notOurTeam3, defaultTeamId, now, false, false, false, 1),
				makeGame(notOurTeam2, defaultTeamId, now, false, false, false, 2),
				makeGame(notOurTeam2, notOurTeam1, now, false, false, false, 3),
			},
		},
		{
			name: "MultipleGamesNoneMatch",
			args: args{
				respSchedule: &schema.Schedule{
					Dates: []schema.Date{
						{
							Games: []schema.Game{
								makeGame(notOurTeam2, notOurTeam1, now, false, false, false, 1),
								makeGame(notOurTeam3, notOurTeam2, now, false, false, false, 2),
								makeGame(notOurTeam1, notOurTeam3, now, false, false, false, 3),
							},
						},
					},
				},
				teamId: defaultTeamId,
			},
			want: []schema.Game{
				makeGame(notOurTeam2, notOurTeam1, now, false, false, false, 1),
				makeGame(notOurTeam3, notOurTeam2, now, false, false, false, 2),
				makeGame(notOurTeam1, notOurTeam3, now, false, false, false, 3),
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
								makeGame(notOurTeam1, defaultTeamId, now, false, false, false, 2),
								makeGame(notOurTeam1, defaultTeamId, past, false, false, false, 1),
							},
						},
					},
				},
				teamId: defaultTeamId,
			},
			want: []schema.Game{
				makeGame(notOurTeam1, defaultTeamId, past, false, false, false, 1),
				makeGame(notOurTeam1, defaultTeamId, now, false, false, false, 2),
			},
		},
		{
			name: "ThreeGamesSameTeamDifferentTimes",
			args: args{
				respSchedule: &schema.Schedule{
					Dates: []schema.Date{
						{
							Games: []schema.Game{
								makeGame(notOurTeam1, defaultTeamId, now, false, false, false, 2),
								makeGame(notOurTeam1, defaultTeamId, future, false, false, false, 3),
								makeGame(notOurTeam1, defaultTeamId, past, false, false, false, 1),
							},
						},
					},
				},
				teamId: defaultTeamId,
			},
			want: []schema.Game{
				makeGame(notOurTeam1, defaultTeamId, past, false, false, false, 1),
				makeGame(notOurTeam1, defaultTeamId, now, false, false, false, 2),
				makeGame(notOurTeam1, defaultTeamId, future, false, false, false, 3),
			},
		},
		{
			name: "ThreeGamesDifferentTimesAndTeams",
			args: args{
				respSchedule: &schema.Schedule{
					Dates: []schema.Date{
						{
							Games: []schema.Game{
								makeGame(notOurTeam2, notOurTeam1, past, false, false, false, 3),
								makeGame(notOurTeam3, defaultTeamId, now, false, false, false, 1),
								makeGame(notOurTeam3, defaultTeamId, future, false, false, false, 2),
							},
						},
					},
				},
				teamId: defaultTeamId,
			},
			want: []schema.Game{
				makeGame(notOurTeam3, defaultTeamId, now, false, false, false, 1),
				makeGame(notOurTeam3, defaultTeamId, future, false, false, false, 2),
				makeGame(notOurTeam2, notOurTeam1, past, false, false, false, 3),
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
								makeGame(notOurTeam1, defaultTeamId, now, true, true, false, 2),
								makeGame(notOurTeam1, defaultTeamId, now, true, false, false, 1),
							},
						},
					},
				},
				teamId: defaultTeamId,
			},
			want: []schema.Game{
				makeGame(notOurTeam1, defaultTeamId, now, true, false, false, 1),
				makeGame(notOurTeam1, defaultTeamId, now, true, true, false, 2),
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
			name: "testing",
			args: args{
				respSchedule: &schema.Schedule{
					Dates: []schema.Date{
						{
							Games: []schema.Game{
								makeGame(notOurTeam1, defaultTeamId, now, true, true, false, 3),
								makeGame(notOurTeam1, defaultTeamId, now, true, false, false, 2),
								makeGame(notOurTeam1, defaultTeamId, future, false, false, true, 1),
							},
						},
					},
				},
				teamId: defaultTeamId,
			},
			want: []schema.Game{
				makeGame(notOurTeam1, defaultTeamId, future, false, false, true, 1),
				makeGame(notOurTeam1, defaultTeamId, now, true, false, false, 2),
				makeGame(notOurTeam1, defaultTeamId, now, true, true, false, 3),
			},
		},

		{
			name: "OneGameLive",
			args: args{
				respSchedule: &schema.Schedule{
					Dates: []schema.Date{
						{
							Games: []schema.Game{
								makeGame(notOurTeam1, defaultTeamId, past, false, false, false, 2),
								makeGame(notOurTeam1, defaultTeamId, now, false, false, true, 1),
								makeGame(notOurTeam1, defaultTeamId, future, false, false, false, 3),
							},
						},
					},
				},
				teamId: defaultTeamId,
			},
			want: []schema.Game{
				makeGame(notOurTeam1, defaultTeamId, now, false, false, true, 1),
				makeGame(notOurTeam1, defaultTeamId, past, false, false, false, 2),
				makeGame(notOurTeam1, defaultTeamId, future, false, false, false, 3),
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
								makeGame(notOurTeam3, notOurTeam2, past, false, false, false, 3),
								makeGame(notOurTeam1, defaultTeamId, now, false, false, false, 1),
								makeGame(notOurTeam1, defaultTeamId, future, false, false, false, 2),
							},
						},
					},
				},
				teamId: defaultTeamId,
			},
			want: []schema.Game{
				makeGame(notOurTeam1, defaultTeamId, now, false, false, false, 1),
				makeGame(notOurTeam1, defaultTeamId, future, false, false, false, 2),
				makeGame(notOurTeam3, notOurTeam2, past, false, false, false, 3),
			},
		},

		{
			name: "TeamsTimesDoubleheaderLive",
			args: args{
				respSchedule: &schema.Schedule{
					Dates: []schema.Date{
						{
							Games: []schema.Game{
								makeGame(defaultTeamId, notOurTeam1, now, true, false, false, 2),
								makeGame(notOurTeam2, notOurTeam1, past, false, true, false, 4),
								makeGame(notOurTeam3, notOurTeam1, future, false, false, false, 5),
								makeGame(notOurTeam1, defaultTeamId, now, true, true, false, 3),
								makeGame(defaultTeamId, notOurTeam1, future, false, false, true, 1),
								makeGame(notOurTeam2, notOurTeam1, past, true, true, false, 6),
								makeGame(notOurTeam3, notOurTeam1, future, false, false, false, 7),
								makeGame(notOurTeam2, notOurTeam1, future, false, true, true, 8),
							},
						},
					},
				},
				teamId: defaultTeamId,
			},
			want: []schema.Game{
				makeGame(defaultTeamId, notOurTeam1, future, false, false, true, 1),
				makeGame(defaultTeamId, notOurTeam1, now, true, false, false, 2),
				makeGame(notOurTeam1, defaultTeamId, now, true, true, false, 3),
				makeGame(notOurTeam2, notOurTeam1, past, false, true, false, 4),
				makeGame(notOurTeam3, notOurTeam1, future, false, false, false, 5),
				makeGame(notOurTeam2, notOurTeam1, past, true, true, false, 6),
				makeGame(notOurTeam3, notOurTeam1, future, false, false, false, 7),
				makeGame(notOurTeam2, notOurTeam1, future, false, true, true, 8),
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
