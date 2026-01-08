package commands

import (
	"fmt"

	"github.com/fc1g/gator/internal/types"
)

func Feeds(state *types.State, _ types.Command) error {
	context, cancel := state.Context()
	defer cancel()

	feeds, err := state.DB.GetFeeds(context)
	if err != nil {
		return fmt.Errorf("error getting feeds: %v", err)
	}

	if len(feeds) == 0 {
		fmt.Println("no feeds found! add one with 'gator addfeed <name> <feed_url>'")
		return nil
	}

	for _, feed := range feeds {
		user, err := state.DB.GetUserById(context, feed.UserID)
		if err != nil {
			return fmt.Errorf("error getting user: %v", err)
		}
		fmt.Println("- ", "name: ", feed.Name, ", ", "url: ", feed.Url, ",", "user: ", user.Name)
	}

	return nil
}
