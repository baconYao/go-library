package itunes

// refs: https://affiliate.itunes.apple.com/resources/documentation/itunes-store-web-service-search-api/

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

// ItunesAPIServices :與 itunes 的 api 溝通的結構
type ItunesAPIServices struct{}

// NewItunesAPIServices :建立一個 ItunesAPIServices instance
func NewItunesAPIServices() *ItunesAPIServices {
	return &ItunesAPIServices{}
}

// SearchResponse :Itunes 的 response
type SearchResponse struct {
	ResultCount int `json:"resultCount"`
	Results     []struct {
		WrapperType            string    `json:"wrapperType"`
		Kind                   string    `json:"kind"`
		CollectionID           int       `json:"collectionId"`
		TrackID                int       `json:"trackId"`
		ArtistName             string    `json:"artistName"`
		CollectionName         string    `json:"collectionName"`
		TrackName              string    `json:"trackName"`
		CollectionCensoredName string    `json:"collectionCensoredName"`
		TrackCensoredName      string    `json:"trackCensoredName"`
		CollectionViewURL      string    `json:"collectionViewUrl"`
		FeedURL                string    `json:"feedUrl,omitempty"`
		TrackViewURL           string    `json:"trackViewUrl"`
		ArtworkURL30           string    `json:"artworkUrl30"`
		ArtworkURL60           string    `json:"artworkUrl60"`
		ArtworkURL100          string    `json:"artworkUrl100"`
		CollectionPrice        float64   `json:"collectionPrice"`
		TrackPrice             float64   `json:"trackPrice"`
		TrackRentalPrice       int       `json:"trackRentalPrice"`
		CollectionHdPrice      int       `json:"collectionHdPrice"`
		TrackHdPrice           int       `json:"trackHdPrice"`
		TrackHdRentalPrice     int       `json:"trackHdRentalPrice"`
		ReleaseDate            time.Time `json:"releaseDate"`
		CollectionExplicitness string    `json:"collectionExplicitness"`
		TrackExplicitness      string    `json:"trackExplicitness"`
		TrackCount             int       `json:"trackCount"`
		Country                string    `json:"country"`
		Currency               string    `json:"currency"`
		PrimaryGenreName       string    `json:"primaryGenreName"`
		ContentAdvisoryRating  string    `json:"contentAdvisoryRating,omitempty"`
		ArtworkURL600          string    `json:"artworkUrl600"`
		GenreIds               []string  `json:"genreIds"`
		Genres                 []string  `json:"genres"`
		ArtistID               int       `json:"artistId,omitempty"`
		ArtistViewURL          string    `json:"artistViewUrl,omitempty"`
	} `json:"results"`
}

// Search :搜尋 itunes' podcast 的 api 介面
func (ias *ItunesAPIServices) Search(term string) (SearchResponse, error) {
	// https://itunes.apple.com/search
	searchURL := url.URL{
		Scheme: "https",
		Host:   "itunes.apple.com",
		Path:   "search",
	}

	q := searchURL.Query()
	q.Set("entity", "podcast")
	q.Set("term", term)

	// entity=podcast&term="Full+Stack+Radio"
	searchURL.RawQuery = q.Encode()

	res, err := http.Get(searchURL.String())
	if err != nil {
		//回傳空的SearchResponse struct & error
		return SearchResponse{}, err
	}

	defer res.Body.Close()

	var searchRespone SearchResponse

	err = json.NewDecoder(res.Body).Decode(&searchRespone)

	return searchRespone, err

}
