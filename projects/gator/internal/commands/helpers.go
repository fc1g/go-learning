package commands

import (
	"strings"

	"github.com/fc1g/gator/internal/types"
)

func ValidateArgs(cmd types.Command, expected int, usage error) error {
	if len(cmd.Args) != expected {
		return usage
	}

	return nil
}

func CleanInput(input string) (cleanedInput string) {
	return strings.Trim(input, " ")
}
