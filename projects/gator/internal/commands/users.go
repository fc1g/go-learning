package commands

import (
	"fmt"

	"github.com/fc1g/gator/internal/types"
)

func Users(state *types.State, _ types.Command) error {
	context, cancel := state.Context()
	defer cancel()

	users, err := state.DB.GetUsers(context)
	if err != nil {
		return fmt.Errorf("error getting users: %v", err)
	}

	if len(users) == 0 {
		fmt.Println("no users found! log in with 'gator login <username>'")
		return nil
	}

	for _, user := range users {
		fmt.Print("- ", user.Name+" ")
		if user.Name == state.Config.CurrentUserName {
			fmt.Println("(current)")
		}
	}

	return nil
}
