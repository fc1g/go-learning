package types

import (
	"context"
	"fmt"
	"time"
)

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	Handlers map[string]func(*State, Command) error
}

func (c *Commands) Run(state *State, command Command) error {
	cmd, ok := c.Handlers[command.Name]
	if !ok {
		return fmt.Errorf("unknown command: %s", command.Name)
	}

	return cmd(state, command)
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	if _, ok := c.Handlers[name]; ok {
		return
	}

	c.Handlers[name] = f
}

func (s *State) Context() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}
