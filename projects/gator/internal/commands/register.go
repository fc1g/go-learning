package commands

import (
	"fmt"
	"time"

	"github.com/fc1g/gator/internal/database"
	"github.com/fc1g/gator/internal/types"
	"github.com/fc1g/gator/pkg/errors"
	"github.com/google/uuid"
)

func Register(state *types.State, command types.Command) error {
	if err := ValidateArgs(command, 1, errors.ErrInvalidRegisterArgsLength); err != nil {
		return err
	}

	username := CleanInput(command.Args[0])

	context, cancel := state.Context()
	defer cancel()

	user, err := state.DB.CreateUser(context, database.CreateUserParams{
		ID:        uuid.New(),
		Name:      username,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}

	err = Login(state, types.Command{Name: "login", Args: []string{username}})
	if err != nil {
		return fmt.Errorf("error logging in user: %v", err)
	}

	fmt.Println("user created successfully:", user)
	return nil
}
