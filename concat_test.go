package gomponents_test

import (
	"fmt"
	"testing"

	g "maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func baseButton(children ...g.Node) g.Node {
	return html.Button(
		html.Class("rounded-md bg-blue-500 py-1 px-2"),
		g.Group(children),
	)
}

func pinkButton(children ...g.Node) g.Node {
	return baseButton(
		html.Class("bg-pink-500"),
		g.Group(children),
	)
}

func TestNoDuplicateClasses(t *testing.T) {
	btn := fmt.Sprintf("%s", pinkButton(
		g.Text("You pink yeah?"),
	))
	expected := `<button class="rounded-md bg-blue-500 py-1 px-2 bg-pink-500">You pink yeah?</button>`
	if btn != expected {
		t.Errorf("DOM mismatch (expected/actual)\n%s\n%s", expected, btn)
	}
}
