package tui

import "fmt"

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func Clamp(v, low, high int) int {
	return Min(Max(v, low), high)
}

func Truncate(s string, limit int) string {
	if len(s) >= Max(limit, 20) {
		return fmt.Sprintf("%v...", s[:limit])
	}
	return s
}
