package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	// dictionary := map[string]string{"test": "this is test"}
	dictionary := Dictionary{"test": "this is test value"}

	t.Run("known word", func(t *testing.T) {
		// got := Search(dictionary, "test")
		got, _ := dictionary.Search("test")
		want := "this is test value"
		assertStrings(t, got, want)

	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknownTest")
		if err == nil {
			t.Fatal("expected to get an error")
		}
		assertStrings(t, err.Error(), ErrNotFound.Error())
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q where given,%q", got, want, "test")
	}
}

// So when you pass a map to a function/method, you are indeed copying it,
// but just the pointer part, not the underlying data structure that contains the data.

func TestAdd(t *testing.T) {
	t.Run("new word ", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "testAdd"
		value := "adding this testcase"
		err := dictionary.Add(word, value)
		assertError(t, err, "nil")
		assertDefination(t, dictionary, word, value)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "testAdd"
		value := "adding this testcase"
		dictionary := Dictionary{word: value}
		err := dictionary.Add(word, "new value")
		assertError(t, err, ErrAlreadyExists.Error())
		assertDefination(t, dictionary, word, value)
	})
}

func assertDefination(t testing.TB, dictionary Dictionary, word, value string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, value)

}

func assertError(t testing.TB, err error, errorString string) {
	// if()
}
