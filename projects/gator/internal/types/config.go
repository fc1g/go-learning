package types

import "github.com/fc1g/gator/internal/database"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

type State struct {
	Config *Config
	DB     *database.Queries
}
