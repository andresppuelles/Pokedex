package main

import "strings"

func cleanInput(text string) []string {
	if len(text) < 1 {
		return []string{}
	}
	cleaned := strings.Fields(text)
	return cleaned
}
