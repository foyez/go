package utils

import "testing"

func TestMakeExcited(t *testing.T) {
	want := "OMG SO EXCITING!"
	got := MakeExcited("omg so exciting")

	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}
