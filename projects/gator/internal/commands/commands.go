package commands

import (
	"github.com/fc1g/gator/internal/middleware"
	"github.com/fc1g/gator/internal/types"
)

func NewCommands() *types.Commands {
	commands := &types.Commands{
		Handlers: make(map[string]func(*types.State, types.Command) error),
	}

	commands.Register("login", Login)
	commands.Register("register", Register)
	commands.Register("reset", Reset)
	commands.Register("users", Users)
	commands.Register("agg", Agg)
	commands.Register("addfeed", middleware.LoggedIn(AddFeed))
	commands.Register("feeds", Feeds)
	commands.Register("follow", middleware.LoggedIn(Follow))
	commands.Register("following", middleware.LoggedIn(Following))
	commands.Register("unfollow", middleware.LoggedIn(Unfollow))
	commands.Register("browse", middleware.LoggedIn(Browse))

	return commands
}
