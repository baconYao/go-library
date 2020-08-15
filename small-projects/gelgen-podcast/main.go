package main

import (
	"fmt"
	"gelgen-podcasr/itunes"
	"log"
)

func main() {
	fmt.Println("Init....")
	ias := itunes.NewItunesAPIServices()

	// 搜尋 itunes 上名為 Full Stacj Radio 的 podcast 資訊
	// res, err := ias.Search("Full Stack Radio")
	res, err := ias.Search("Full Stack Radio")
	if err != nil {
		log.Fatal("Error while serching: %v", err)
	}

	for _, item := range res.Results {
		log.Println("---------------------")
		log.Printf("Artist: %s", item.ArtistName)
		log.Printf("Podcast Name: %s", item.TrackName)
		log.Printf("Feed URL: %s", item.FeedURL)
		log.Println("---------------------")
	}
}
