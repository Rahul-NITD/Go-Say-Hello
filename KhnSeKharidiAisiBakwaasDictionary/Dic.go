package khnsekharidiaisibakwaasdictionary

import "errors"

type Dictionary map[string]string

func (dictionary Dictionary) Search(searchWord string) (string, error) {
	result, ok := dictionary[searchWord]
	if !ok {
		return result, errors.New("search word not found in Dictionary")
	}
	return result, nil
}
