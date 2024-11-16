package racer

import "testing"

func TestRacer(t *testing.T) {
	slowURL := "vk.com"
	fastURL := "monkeytype.com"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
