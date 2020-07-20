package main

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/baconYao/go-library/small-projects/nhl-project/nhlapi"
)

func main() {
	// help benchmarking the request time
	now := time.Now()

	rosterFile, err := os.OpenFile("rosters.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
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

	for _, team := range teams {
		log.Println("----------------------")
		log.Println("Name:", team.Name)
		log.Println("----------------------")
	}

	log.Printf("took %v", time.Now().Sub(now).String())
}
