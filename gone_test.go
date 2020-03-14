package gone

import "testing"

func TestGone(t *testing.T) {
	want := "Hello, world!"
	if got := Gone(); got != want {
		t.Errorf("Gone() = %q, want %q", got, want)
	}
}
