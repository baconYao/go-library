package main

import (
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/baconYao/go-library/small-projects/nhl-project/nhlapi"
)

func main() {
	// help benchmarking the request time
	now := time.Now()

	rosterFile, err := os.OpenFile("rosters.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal("Error opening the file rosters.txt")
		log.Fatal(err)
	}
	defer rosterFile.Close()

	wrt := io.MultiWriter(os.Stdout, rosterFile)

	log.SetOutput(wrt)

	teams, err := nhlapi.GetAllTeams()
	if err != nil {
		log.Fatalf("Error while getting all teams :%v", err)
	}

	// In order to use goroutine and wait response, so here use WaitGroup to wait all goroutines finished.
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(teams))

	// Create unbuffered channel
	results := make(chan []nhlapi.Roster)

	for _, team := range teams {
		// Pass each team's id and get back its roster
		go func(team nhlapi.Team) {
			roster, err := nhlapi.GetRosters(team.ID)
			if err != nil {
				log.Fatalf("Error getting roster: %v", err)
			}

			results <- roster

			waitGroup.Done()
		}(team)
	}

	go func() {
		waitGroup.Wait()
		// Close the channel
		close(results)
	}()

	// Show all players in rosters
	display(results)

	log.Printf("took %v", time.Now().Sub(now).String())
}

func display(results chan []nhlapi.Roster) {
	for result := range results {
		for _, ros := range result {
			log.Println("---------------------------")
			log.Printf("ID: %d\n", ros.Person.ID)
			log.Printf("Name: %s\n", ros.Person.FullName)
			log.Printf("Position: %s\n", ros.Position.Abbreviation)
			log.Printf("Jersey: %s\n", ros.JerseyNumber)
			log.Println("---------------------------")
		}
	}
}
