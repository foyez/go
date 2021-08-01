package hello

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	assertErrorMessage := func(t testing.TB, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Foyez")
		want := "Hello, Foyez"
		assertErrorMessage(t, got, want)
	})
}

func ExampleHello() {
	greeting := Hello("Zayan")
	fmt.Println(greeting)
	// Output: Hello, Zayan
}
