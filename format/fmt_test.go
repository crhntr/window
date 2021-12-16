package format

import (
	"testing"
)

func TestSprintf(t *testing.T) {
	match := func(t *testing.T, got, exp string) {
		t.Helper()
		if got != exp {
			t.Errorf("expected %q but got %q", exp, got)
		}
	}

	match(t,
		Sprintf("<div>%d</div>", 10),
		"<div>10</div>",
	)
	match(t,
		Sprintf("<div>%s</div>", "<br/>"),
		"<div>&lt;br/&gt;</div>",
	)
	match(t,
		Sprintf("<div>%q</div>", "<br/>"),
		"<div>&#34;&lt;br/&gt;&#34;</div>",
	)
	match(t,
		Sprintf(`<div style="width: %[2]dpx;">%[1]s</div>`, "<br/>", 20),
		`<div style="width: 20px;">&lt;br/&gt;</div>`,
	)
}
