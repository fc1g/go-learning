package commands

import (
	"fmt"

	"github.com/fc1g/gator/internal/types"
)

func Reset(state *types.State, _ types.Command) error {
	context, cancel := state.Context()
	defer cancel()

	err := state.DB.DeleteAllUsers(context)
	if err != nil {
		return fmt.Errorf("error deleting all users: %v", err)
	}

	return nil
}
