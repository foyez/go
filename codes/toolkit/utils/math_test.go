package utils

import "testing"

func TestAdd(t *testing.T) {
	want := 6
	got := Add(1, 2, 3)

	if want != got {
		t.Errorf("want %d got %d", want, got)
	}
}
