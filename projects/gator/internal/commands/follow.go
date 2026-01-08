package commands

import (
	"fmt"
	"time"

	"github.com/fc1g/gator/internal/database"
	"github.com/fc1g/gator/internal/types"
	"github.com/fc1g/gator/pkg/errors"
	"github.com/google/uuid"
)

func Follow(state *types.State, command types.Command, user database.User) error {
	if len(command.Args) != 1 {
		return fmt.Errorf("usage: gator follow <feed_url>")
	}
	if err := ValidateArgs(command, 1, errors.ErrInvalidFollowArgsLength); err != nil {
		return err
	}

	context, cancel := state.Context()
	defer cancel()

	feed, err := state.DB.GetFeedByUrl(context, command.Args[0])
	if err != nil {
		return fmt.Errorf("error getting feed: %v", err)
	}

	follow, err := state.DB.CreateFeedFollow(context, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("error creating feed follow: %v", err)
	}

	fmt.Println("feed followed successfully:", follow)

	return nil
}
