package commands

import (
	"fmt"
	"strings"
	"time"

	"github.com/fc1g/gator/internal/database"
	"github.com/fc1g/gator/internal/rss"
	"github.com/fc1g/gator/internal/types"
	"github.com/fc1g/gator/pkg/errors"
	"github.com/google/uuid"
)

func parsePublishedAt(pubDate string) time.Time {
	formats := []string{
		time.RFC1123Z,
		time.RFC1123,
		time.RFC822Z,
		time.RFC822,
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02 15:04:05",
		"Mon, 02 Jan 2006 15:04:05 -0700",
	}

	for _, format := range formats {
		if t, err := time.Parse(format, pubDate); err == nil {
			return t
		}
	}

	return time.Now()
}

func scrapeFeeds(state *types.State) error {
	context, cancel := state.Context()
	defer cancel()

	feed, err := state.DB.GetNextFeedToFetch(context)
	if err != nil {
		return fmt.Errorf("couldn't get next feed to fetch: %v", err)
	}

	fmt.Printf("Fetching feed: %s (URL: %s)\n", feed.Name, feed.Url)

	_, err = state.DB.MarkFeedFetched(context, feed.ID)
	if err != nil {
		return fmt.Errorf("couldn't mark feed as fetched: %v", err)
	}

	rssFeed, err := rss.FetchFeed(context, feed.Url)
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %v", err)
	}

	for _, item := range rssFeed.Channel.Item {
		publishedAt := parsePublishedAt(item.PubDate)

		_, err := state.DB.CreatePost(context, database.CreatePostParams{
			ID:          uuid.New(),
			Title:       item.Title,
			Description: item.Description,
			Url:         item.Link,
			FeedID:      feed.ID,
			PublishedAt: publishedAt,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") ||
				strings.Contains(err.Error(), "UNIQUE constraint failed") {
				continue
			}
			fmt.Printf("Error saving post '%s': %v\n", item.Title, err)
			continue
		}
	}

	return nil
}

func Agg(state *types.State, command types.Command) error {
	if err := ValidateArgs(command, 1, errors.ErrInvalidAggArgsLength); err != nil {
		return err
	}

	timeBetweenRequests, err := time.ParseDuration(command.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	fmt.Printf("Collecting feeds every %s\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	defer ticker.Stop()

	for ; ; <-ticker.C {
		err := scrapeFeeds(state)
		if err != nil {
			fmt.Printf("Error scraping feeds: %v\n", err)
			continue
		}
	}
}
