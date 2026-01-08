package commands

import (
	"fmt"

	"github.com/fc1g/gator/internal/database"
	"github.com/fc1g/gator/internal/types"
	"github.com/fc1g/gator/pkg/errors"
)

func Unfollow(state *types.State, command types.Command, user database.User) error {
	if err := ValidateArgs(command, 1, errors.ErrInvalidUnfollowArgsLength); err != nil {
		return err
	}

	context, cancel := state.Context()
	defer cancel()

	feed, err := state.DB.GetFeedByUrl(context, command.Args[0])
	if err != nil {
		return fmt.Errorf("error getting feed: %v", err)
	}

	err = state.DB.DeleteFeedFollow(context, database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error deleting feed follow: %v", err)
	}

	fmt.Println("feed unfollowed successfully:", feed)

	return nil
}
