package main

import (
	"MLBTakeHome/sortgames"
	"flag"
	"fmt"
	"log"
	"os"
)

// main takes in inputs for date and teamId
// After calling sorting function, main either writes the newly sorted JSON to stdout or writes it to
// the file 'schedule.go'
func main() {
	dateFlag := flag.String("date", "2016-11-02", "date in format YYYY-MM-DD")
	teamFlag := flag.Int("teamId", 112, "Id of favorite team, e.g. 112 for CHC")
	asFileFlag := flag.Bool("asFile", false, "if this flag is set to true then the return value will be written to the file 'schedule.json'")

	flag.Parse()

	resp, err := sortgames.FavoriteTeamGames(*teamFlag, *dateFlag)

	if err != nil {
		err = os.WriteFile("error.txt", []byte(err.Error()), 0666)
		if err != nil {
			log.Fatal(err)
		}
	}

	if *asFileFlag {
		fmt.Println("Generating sorted schedule...")
		err = os.WriteFile("schedule.json", resp, 0666)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Generated schedule in 'schedule.json'")
	} else {
		fmt.Println(string(resp))
	}
}
