package main

import "testing"

// func TestHello(t *testing.T) {
// 	got := Hello("vaibhava")
// 	want := "Hello, vaibhava"

// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "english")
		want := "Hello, Chris"

		assertCorrectResults(t, got, want)

	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", french)
		want := "Bonjour, World"

		assertCorrectResults(t, got, want)
	})
	t.Run("greeting in Spanish", func(t *testing.T) {
		got := Hello("Neymar", "Spanish")
		want := "Hola amigo, Neymar"

		assertCorrectResults(t, got, want)
	})
}

func assertCorrectResults(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
