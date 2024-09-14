package main

func pointTo[T any](value T) *T {
	return &value
}
