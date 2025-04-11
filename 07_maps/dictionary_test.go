package main

import "testing"

func assertDefinition(t testing.TB, d Dictionary, word, definition string) {
	t.Helper()

	got, err := d.Search(word)
	if err != nil {
		t.Fatal("should find the added word:", err)
	}
	assertStrings(t, got, definition)
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a string"}

	t.Run("known word", func(t *testing.T) {
		want := "this is just a string"
		assertDefinition(t, dictionary, "test", want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		if got == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		err := dictionary.Add(word, "this is just a test")
		want := "this is just a test"

		assertDefinition(t, dictionary, word, want)
		assertNoError(t,err)
	})

	t.Run("existing word", func(t *testing.T) {
		testWord := "test word"
		testDefinition := "test definition"

		dictionary := Dictionary{}
		err:= dictionary.Add(testWord, testDefinition)
		assertNoError(t,err)

		err = dictionary.Add(testWord, "new definition")

		assertError(t, err, ErrWordExists)
		//assertDefinition(t, dictionary, testWord, testDefinition)

	})

}

func TestUpdate (t *testing.T){
	t.Run("existing word", func (t *testing.T){
		word:="test"
		defnt:= "test definition"
		dictionary := Dictionary{}
		dictionary.Add(word,defnt)
	
		newDefns := "new definition"
		dictionary.Update(word,newDefns)
		assertDefinition(t, dictionary, word,newDefns)
	})

	t.Run("new word", func(t *testing.T){
		dictionary := Dictionary{}

		word:="test word"
		testDef :="test def"

		err:=dictionary.Update(word,testDef)
		assertError(t, err,ErrWordDoesNotExists)
	})
	
}

func TestDelete (t *testing.T){
	t.Run("existing word", func (t *testing.T){
		word:="test word"
		def := "test def"
		dictionary := Dictionary{}

		dictionary.Add(word,def)

		err := dictionary.Delete(word)

		assertNoError(t, err)
		_,err = dictionary.Search(word)
		assertError(t, err, ErrNotFound)

	})

	t.Run("non existing word", func (t *testing.T){
		dictionary := Dictionary{}

		err := dictionary.Delete("test")
		assertError(t,err,ErrWordDoesNotExists)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q but want %q", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	if err != nil {
		t.Fatal("Expected no error, but got an error")
	}
}
