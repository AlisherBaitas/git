// Command main prints a health report for the git history of the current repository.
package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/AlisherBaitas/git/internal/health"
)

func main() {
	out, err := exec.Command("git", "log", "--pretty=format:%ad|%s", "--date=short").Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, "cannot read git log:", err)
		os.Exit(1)
	}

	var first, last time.Time
	commits := 0
	total := 0

	for _, line := range strings.Split(strings.TrimSpace(string(out)), "\n") {
		parts := strings.SplitN(line, "|", 2)
		if len(parts) != 2 {
			continue
		}
		date, err := time.Parse("2006-01-02", parts[0])
		if err != nil {
			continue
		}
		if commits == 0 || date.Before(first) {
			first = date
		}
		if commits == 0 || date.After(last) {
			last = date
		}
		commits++
		total += health.ScoreMessage(parts[1])
	}

	if commits == 0 {
		fmt.Println("No commits found.")
		return
	}

	days := int(last.Sub(first).Hours() / 24)
	span := max(1, days)
	cadence := float64(commits) / float64(span)
	avg := float64(total) / float64(commits)

	fmt.Println("=== Repo Health ===")
	fmt.Printf("Commits: %d\n", commits)
	fmt.Printf("Span: %d days\n", span)
	fmt.Printf("Cadence: %.2f commits/day\n", cadence)
	fmt.Printf("Average message score: %.1f / 3\n", avg)
	fmt.Printf("History: %s\n", health.Label(avg))
}
