package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"

	"github.com/fc1g/gator/internal/types"
)

func FetchFeed(ctx context.Context, feedURL string) (*types.RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &types.RSSFeed{}, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("User-Agent", "gator")
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return &types.RSSFeed{}, fmt.Errorf("error fetching feed: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("error closing response body:", err)
		}
	}(res.Body)

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return &types.RSSFeed{}, fmt.Errorf("error reading response body: %v", err)
	}

	var feed types.RSSFeed
	if err := xml.Unmarshal(bytes, &feed); err != nil {
		return &types.RSSFeed{}, fmt.Errorf("error parsing XML: %v", err)
	}

	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)

	for i := range feed.Channel.Item {
		feed.Channel.Item[i].Title = html.UnescapeString(feed.Channel.Item[i].Title)
		feed.Channel.Item[i].Description = html.UnescapeString(feed.Channel.Item[i].Description)
	}

	return &feed, nil
}
