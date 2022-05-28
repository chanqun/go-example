package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("not Found")
var errAlreadyExists = errors.New("word already exists")
var errCanUpdate = errors.New("word does not exists")

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

// Update the word
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)

	if err == errNotFound {
		return errCanUpdate
	}

	d[word] = def

	return nil
}

// Delete the word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
