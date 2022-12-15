package sortgames

import (
	"MLBTakeHome/schema"
	"testing"
)

func TestGameSorter_byFavoriteTeam(t *testing.T) {
	teamId := 133
	type fields struct {
		games           []schema.Game
		comparisonFuncs []comparisonFunc
		teamId          int
	}
	type args struct {
		game1 *schema.Game
		game2 *schema.Game
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "NoGames",
			args: args{
				game1: nil,
				game2: nil,
			},
			want: false,
		},
		{
			name: "FavTeamFirst",
			args: args{
				game1: &schema.Game{
					Teams: schema.Teams{
						Away: schema.Away{
							Team: schema.Team{
								Id: teamId,
							},
						},
						Home: schema.Home{
							Team: schema.Team{
								Id: 0,
							},
						},
					},
				},
				game2: &schema.Game{
					Teams: schema.Teams{
						Away: schema.Away{
							Team: schema.Team{
								Id: 0,
							},
						},
						Home: schema.Home{
							Team: schema.Team{
								Id: 0,
							},
						},
					},
				},
			},
			want: true,
		},
		{
			name: "FavTeamSecond",
			args: args{
				game1: &schema.Game{
					Teams: schema.Teams{
						Away: schema.Away{
							Team: schema.Team{
								Id: 0,
							},
						},
						Home: schema.Home{
							Team: schema.Team{
								Id: 0,
							},
						},
					},
				},
				game2: &schema.Game{
					Teams: schema.Teams{
						Away: schema.Away{
							Team: schema.Team{
								Id: teamId,
							},
						},
						Home: schema.Home{
							Team: schema.Team{
								Id: 0,
							},
						},
					},
				},
			},
			want: false,
		},
		{
			name: "BothTeamsFav",
			args: args{
				game1: &schema.Game{
					Teams: schema.Teams{
						Away: schema.Away{
							Team: schema.Team{
								Id: teamId,
							},
						},
						Home: schema.Home{
							Team: schema.Team{
								Id: 0,
							},
						},
					},
				},
				game2: &schema.Game{
					Teams: schema.Teams{
						Away: schema.Away{
							Team: schema.Team{
								Id: teamId,
							},
						},
						Home: schema.Home{
							Team: schema.Team{
								Id: 0,
							},
						},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gs := &GameSorter{
				teamId: teamId,
			}
			if got := gs.byFavoriteTeam(tt.args.game1, tt.args.game2); got != tt.want {
				t.Errorf("byFavoriteTeam() = %v, want %v", got, tt.want)
			}
		})
	}
}
