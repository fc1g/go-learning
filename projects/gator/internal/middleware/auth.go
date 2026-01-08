package middleware

import (
	"context"
	"fmt"

	"github.com/fc1g/gator/internal/database"
	"github.com/fc1g/gator/internal/types"
	"github.com/fc1g/gator/pkg/errors"
)

func LoggedIn(handler func(*types.State, types.Command, database.User) error) func(*types.State, types.Command) error {
	return func(state *types.State, cmd types.Command) error {
		if state.Config.CurrentUserName == "" {
			return errors.ErrNotLoggedIn
		}

		currentUser, err := state.DB.GetUser(context.Background(), state.Config.CurrentUserName)
		if err != nil {
			return fmt.Errorf("error getting current user: %v", err)
		}

		return handler(state, cmd, currentUser)
	}
}
