package main

import "strings"

func pointTo[T any](value T) *T {
	return &value
}

func removeSpaces(value string) string {
	return strings.ReplaceAll(value, " ", "")
}
