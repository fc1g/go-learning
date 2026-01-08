package commands

import (
	"fmt"
	"strconv"

	"github.com/fc1g/gator/internal/database"
	"github.com/fc1g/gator/internal/types"
)

func Browse(state *types.State, command types.Command, user database.User) error {
	limit := 2

	if len(command.Args) > 0 {
		parsedLimit, err := strconv.Atoi(command.Args[0])
		if err != nil {
			return fmt.Errorf("invalid limit: %v", err)
		}
		limit = parsedLimit
	}

	context, cancel := state.Context()
	defer cancel()

	posts, err := state.DB.GetPostsForUser(context, database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("error getting posts: %v", err)
	}

	if len(posts) == 0 {
		fmt.Println("no posts found! follow some feeds first!")
		return nil
	}

	for _, post := range posts {
		fmt.Println("-", "title:", post.Title, "description:", post.Description, "url:", post.Url)
	}

	return nil
}
