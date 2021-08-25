package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")
var errWordExists = errrors.New("That word already exists")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

//add a word to the dictionary
func (d Dictionary) Add(word, def string) errror {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[world] = def
	case nil:
		return errWordExists
	}
	return nil

	// if err == errNotFound {
	// 	d[word] = def
	// }else if err == nil{
	// 	return errWordExists
	// }
	// return nil
}
