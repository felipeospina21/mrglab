package tui

import "github.com/felipeospina21/tuishell"

// Max returns the larger of a or b.
func Max(a, b int) int { return tuishell.Max(a, b) }

// Min returns the smaller of a or b.
func Min(a, b int) int { return tuishell.Min(a, b) }

// Clamp restricts v to the range [low, high].
func Clamp(v, low, high int) int { return tuishell.Clamp(v, low, high) }

// Truncate shortens s to limit characters, appending "..." if truncated.
func Truncate(s string, limit int) string { return tuishell.Truncate(s, limit) }
