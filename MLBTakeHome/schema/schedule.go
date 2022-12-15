package schema

import (
	"time"
)

type Schedule struct {
	Copyright            string `json:"copyright"`
	TotalItems           int    `json:"totalItems"`
	TotalEvents          int    `json:"totalEvents"`
	TotalGames           int    `json:"totalGames"`
	TotalGamesInProgress int    `json:"totalGamesInProgress"`
	Dates                []Date `json:"dates"`
}

type Date struct {
	Date                 string        `json:"date"`
	TotalItems           int           `json:"totalItems"`
	TotalEvents          int           `json:"totalEvents"`
	TotalGames           int           `json:"totalGames"`
	TotalGamesInProgress int           `json:"totalGamesInProgress"`
	Games                []Game        `json:"games"`
	Events               []interface{} `json:"events"`
}

type Game struct {
	GamePk       int       `json:"gamePk"`
	Link         string    `json:"link"`
	GameType     string    `json:"gameType"`
	Season       string    `json:"season"`
	GameDate     time.Time `json:"gameDate"`
	OfficialDate string    `json:"officialDate"`
	Status       Status    `json:"status"`
	Teams        Teams     `json:"teams"`
	Venue        struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"venue"`
	Content struct {
		Link string `json:"link"`
	} `json:"content"`
	IsTie                  bool   `json:"isTie"`
	GameNumber             int    `json:"gameNumber"`
	PublicFacing           bool   `json:"publicFacing"`
	DoubleHeader           string `json:"doubleHeader"`
	GamedayType            string `json:"gamedayType"`
	Tiebreaker             string `json:"tiebreaker"`
	CalendarEventID        string `json:"calendarEventID"`
	SeasonDisplay          string `json:"seasonDisplay"`
	DayNight               string `json:"dayNight"`
	ScheduledInnings       int    `json:"scheduledInnings"`
	ReverseHomeAwayStatus  bool   `json:"reverseHomeAwayStatus"`
	InningBreakLength      int    `json:"inningBreakLength,omitempty"`
	GamesInSeries          int    `json:"gamesInSeries"`
	SeriesGameNumber       int    `json:"seriesGameNumber"`
	SeriesDescription      string `json:"seriesDescription"`
	RecordSource           string `json:"recordSource"`
	IfNecessary            string `json:"ifNecessary"`
	IfNecessaryDescription string `json:"ifNecessaryDescription"`
}

type Teams struct {
	Away Away `json:"away"`
	Home Home `json:"home"`
}

type Away struct {
	LeagueRecord struct {
		Wins   int    `json:"wins"`
		Losses int    `json:"losses"`
		Pct    string `json:"pct"`
	} `json:"leagueRecord"`
	Score        int  `json:"score"`
	Team         Team `json:"team"`
	IsWinner     bool `json:"isWinner"`
	SplitSquad   bool `json:"splitSquad"`
	SeriesNumber int  `json:"seriesNumber"`
}

type Home struct {
	LeagueRecord struct {
		Wins   int    `json:"wins"`
		Losses int    `json:"losses"`
		Pct    string `json:"pct"`
	} `json:"leagueRecord"`
	Score        int  `json:"score"`
	Team         Team `json:"team"`
	IsWinner     bool `json:"isWinner"`
	SplitSquad   bool `json:"splitSquad"`
	SeriesNumber int  `json:"seriesNumber"`
}

type Team struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}

type Status struct {
	AbstractGameState string `json:"abstractGameState"`
	CodedGameState    string `json:"codedGameState"`
	DetailedState     string `json:"detailedState"`
	StatusCode        string `json:"statusCode"`
	StartTimeTBD      bool   `json:"startTimeTBD"`
	AbstractGameCode  string `json:"abstractGameCode"`
}
