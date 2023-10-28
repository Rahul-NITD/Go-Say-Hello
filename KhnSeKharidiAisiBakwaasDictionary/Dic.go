package khnsekharidiaisibakwaasdictionary

import "errors"

type Dictionary map[string]string

var ErrSearchWordNotFound = errors.New("search word not found in Dictionary")

func (dictionary Dictionary) Search(searchWord string) (string, error) {
	result, ok := dictionary[searchWord]
	if !ok {
		return result, ErrSearchWordNotFound
	}
	return result, nil
}
