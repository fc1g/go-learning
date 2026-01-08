package commands

import (
	"fmt"

	"github.com/fc1g/gator/internal/database"
	"github.com/fc1g/gator/internal/types"
)

func Following(state *types.State, _ types.Command, user database.User) error {
	context, cancel := state.Context()
	defer cancel()

	following, err := state.DB.GetFeedFollowForUser(context, user.ID)
	if err != nil {
		return fmt.Errorf("error getting following feeds: %v", err)
	}

	for _, feed := range following {
		fmt.Println("-", feed.FeedName)
	}
	return nil
}
