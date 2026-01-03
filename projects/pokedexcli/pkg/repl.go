package pkg

import "strings"

func CleanInput(text string) (cleanedText []string) {
	for _, field := range strings.Fields(text) {
		cleanedText = append(cleanedText, strings.ToLower(strings.Trim(field, "")))
	}

	return
}
