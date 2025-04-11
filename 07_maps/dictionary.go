package main

type Dictionary map[string]string

const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("word already exists")
	ErrWordDoesNotExists = DictionaryErr("word does not exist")
)

type DictionaryErr string 

func (e DictionaryErr) Error() string {
	return string (e)
}

func (d Dictionary) Search(key string) (string, error) {

	definition, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(key string, definition string) error {

	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = definition
	case nil: 
		return ErrWordExists
	}
	
	return nil

}

func (d Dictionary) Update (word string, definition string)error {
	_,err:= d.Search(word)
	switch err {
	case ErrNotFound: 
		return ErrWordDoesNotExists
	case nil: 
		d[word] = definition
	}
	return nil
}

func (d Dictionary) Delete (word string) error {

	_,err := d.Search(word)
	
	switch err {
	case nil: 
		delete(d,word)
		return nil
	case ErrNotFound: 
		return ErrWordDoesNotExists
	}
	return nil
}
