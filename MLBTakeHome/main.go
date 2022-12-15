package main

import (
	"MLBTakeHome/sortgames"
	"flag"
	"fmt"
	"log"
	"os"
)

// main takes in inputs for date and teamId
// After calling sorting function, main then writes the newly sorted JSON to
// the file schedule.go
// Currently drops returned http errors, and will simply return an empty file
func main() {
	dateFlag := flag.String("date", "2016-11-02", "date in format YYYY-MM-DD")
	teamFlag := flag.Int("teamId", 112, "Id of favorite team, e.g. 112 for CHC")

	flag.Parse()

	resp, err := sortgames.FavoriteTeamGames(*teamFlag, *dateFlag)

	if err != nil {
		err = os.WriteFile("error.txt", []byte(err.Error()), 0666)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("TeamId: %v\n", *teamFlag)
	fmt.Printf("Date: %v\n", *dateFlag)
	fmt.Println("Generating sorted schedule...")
	err = os.WriteFile("schedule.json", resp, 0666)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Generated schedule in 'schedule.json'")
}
