package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("not Found")
var errAlreadyExists = errors.New("word already exists")

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]

	if exists {
		return value, nil
	}

	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	if err != errNotFound {
		return errAlreadyExists
	}

	d[word] = def

	return nil
}

func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)

	if err == errNotFound {
		return err
	}

	d[word] = def

	return nil
}
