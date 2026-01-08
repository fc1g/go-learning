package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/fc1g/gator/internal/commands"
	"github.com/fc1g/gator/internal/config"
	"github.com/fc1g/gator/internal/database"
	"github.com/fc1g/gator/internal/types"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	state := types.State{
		Config: cfg,
	}

	db, err := sql.Open("postgres", state.Config.DbURL)
	if err != nil {
		fmt.Println("error opening database:", err)
		os.Exit(1)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		fmt.Println("error connecting to database:", err)
		os.Exit(1)
	}

	state.DB = database.New(db)

	cmds := commands.NewCommands()

	if len(os.Args) < 2 {
		fmt.Println("usage: gator <command> [args...]")
		os.Exit(1)
	}

	err = cmds.Run(&state, types.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
