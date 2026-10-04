// Package health scores commit messages and labels a commit history.
package health

import "strings"

// typeVerbs are the allowed commit types from the project convention.
var typeVerbs = []string{"add", "update", "fix", "delete", "docs", "refactor"}

// topics are words that count as a concrete topic of this project.
var topics = []string{"branch", "merge", "conflict", "remote", "analyzer", "readme", "workflow", "feature"}

// ScoreMessage returns 0..3 points for a commit message:
// +1 for 3 or more words, +1 for starting with a type verb,
// +1 for referencing a file or a concrete topic.
func ScoreMessage(msg string) int {
	words := strings.Fields(msg)
	score := 0

	if len(words) >= 3 {
		score++
	}
	if len(words) > 0 && startsWithType(words[0]) {
		score++
	}
	if referencesTopic(words) {
		score++
	}
	return score
}

// startsWithType reports whether the first word is a type verb, with or without a trailing colon.
func startsWithType(first string) bool {
	first = strings.ToLower(strings.TrimSuffix(first, ":"))
	for _, v := range typeVerbs {
		if first == v {
			return true
		}
	}
	return false
}

// referencesTopic reports whether any word is a file name, an all-caps name like README, or a known topic.
func referencesTopic(words []string) bool {
	for _, w := range words {
		w = strings.Trim(w, ",.:;()")
		if strings.Contains(w, ".") {
			return true
		}
		if len(w) >= 3 && w == strings.ToUpper(w) && w != strings.ToLower(w) {
			return true
		}
		lower := strings.ToLower(w)
		for _, t := range topics {
			if lower == t {
				return true
			}
		}
	}
	return false
}

// Label maps an average message score to a history label.
func Label(avg float64) string {
	switch {
	case avg >= 2.5:
		return "tidy"
	case avg >= 1.5:
		return "acceptable"
	default:
		return "messy"
	}
}
