package commands

import (
	"context"
	"fmt"

	"github.com/fc1g/gator/internal/config"
	"github.com/fc1g/gator/internal/types"
	"github.com/fc1g/gator/pkg/errors"
)

func Login(state *types.State, command types.Command) error {
	if err := ValidateArgs(command, 1, errors.ErrInvalidLoginArgsLength); err != nil {
		return err
	}

	userName := CleanInput(command.Args[0])

	user, err := state.DB.GetUser(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("error getting user: %v", err)
	}

	state.Config.CurrentUserName = userName
	fmt.Println("the user has been set to:", user)

	return config.Write(state.Config)
}
