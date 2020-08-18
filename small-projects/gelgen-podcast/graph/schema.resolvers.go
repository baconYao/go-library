package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"gelgen-podcasr/feeds"
	"gelgen-podcasr/graph/generated"
	"gelgen-podcasr/graph/model"
	"gelgen-podcasr/itunes"
	"gelgen-podcasr/utils"
)

// Query - search
func (r *queryResolver) Search(ctx context.Context, term string) ([]*model.PodCast, error) {
	ias := itunes.NewItunesAPIServices()

	// 搜尋 itunes 上名為 Full Stacj Radio 的 podcast 資訊
	// res, err := ias.Search("Full Stack Radio")
	res, err := ias.Search("Full Stack Radio")
	if err != nil {
		return nil, err
	}

	var podcasts []*model.PodCast

	for _, res := range res.Results {
		podcast := &model.PodCast{
			Artist:        res.ArtistName,
			PodcastName:   res.TrackName,
			FeedURL:       res.FeedURL,
			Thumbnail:     res.ArtworkURL100,
			EpisodesCount: res.TrackCount,
			Genres:        res.Genres,
		}

		podcasts = append(podcasts, podcast)
	}

	return podcasts, nil
}

// Query - feed
func (r *queryResolver) Feed(ctx context.Context, feedURL string) ([]*model.FeedItem, error) {
	res, err := feeds.GetFeed(feedURL)
	if err != nil {
		return nil, err
	}

	var feedItems []*model.FeedItem

	for _, item := range res.Channel.Item {
		feedItem := &model.FeedItem{
			PubDate:     item.PubDate,
			Text:        item.Text,
			Title:       item.Title,
			Subtitle:    item.Subtitle,
			Description: item.Description,
			Image:       utils.CheckNullString(item.Image.Href),
			Summary:     item.Summary,
			LinkURL:     item.Enclosure.URL,
			Duration:    item.Duration,
		}

		feedItems = append(feedItems, feedItem)
	}

	return feedItems, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
