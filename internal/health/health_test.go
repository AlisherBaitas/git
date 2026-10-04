package health

import "testing"

func TestScoreMessage(t *testing.T) {
	cases := []struct {
		msg  string
		want int
	}{
		{"add: parser for stack input in parser.go", 3},
		{"update: README description on main branch", 3},
		{"fix: resolve README merge conflict", 3},
		{"docs: add workflow.md describing branches", 3},
		{"Add: README with description", 3},
		{"update README", 2},
		{"changed some stuff today", 1},
		{"fix typo", 1},
		{"wip", 0},
		{"addition: something", 0},
		{"", 0},
		{"fix: typo in README, again", 3},
		{"update: (workflow) section text", 3},
		{"fix: 123 456", 2},
	}
	for _, c := range cases {
		if got := ScoreMessage(c.msg); got != c.want {
			t.Errorf("ScoreMessage(%q) = %d, want %d", c.msg, got, c.want)
		}
	}
}

func TestLabel(t *testing.T) {
	cases := []struct {
		avg  float64
		want string
	}{
		{3.0, "tidy"},
		{2.6, "tidy"},
		{2.5, "tidy"},
		{2.4, "acceptable"},
		{1.5, "acceptable"},
		{1.49, "messy"},
		{0, "messy"},
	}
	for _, c := range cases {
		if got := Label(c.avg); got != c.want {
			t.Errorf("Label(%v) = %q, want %q", c.avg, got, c.want)
		}
	}
}
