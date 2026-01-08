package commands

import (
	"fmt"
	"time"

	"github.com/fc1g/gator/internal/database"
	"github.com/fc1g/gator/internal/types"
	"github.com/fc1g/gator/pkg/errors"
	"github.com/google/uuid"
)

func AddFeed(state *types.State, command types.Command, user database.User) error {
	if err := ValidateArgs(command, 2, errors.ErrInvalidAddFeedArgsLength); err != nil {
		return err
	}

	context, cancel := state.Context()
	defer cancel()

	feed, err := state.DB.AddFeed(context, database.AddFeedParams{
		ID:        uuid.New(),
		Name:      command.Args[0],
		Url:       command.Args[1],
		UserID:    user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("error adding feed: %v", err)
	}

	fmt.Println("feed added successfully:", feed)

	_, err = state.DB.CreateFeedFollow(context, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("error creating feed follow: %v", err)
	}

	return nil
}
