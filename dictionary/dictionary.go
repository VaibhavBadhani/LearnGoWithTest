package dictionary

import "errors"

var ErrNotFound = errors.New("could not find the word you were looking")
var ErrAlreadyExists = errors.New("the word already exists in the dictionary")

func Search(dictionary Dictionary, key string) string {
	return dictionary[key]
}

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	val, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return val, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrAlreadyExists
	default:
		return err
	}

	return nil
}
