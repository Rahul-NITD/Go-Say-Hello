package khnsekharidiaisibakwaasdictionary

type Dictionary map[string]string

const (
	ErrSearchWordNotFound         = DictionaryErr("search word not found in Dictionary")
	ErrKeyAlreadyExist            = DictionaryErr("key already exists")
	ErrCannotUpdateNonExistingKey = DictionaryErr("cannot update non existing key")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (dictionary Dictionary) Search(searchWord string) (string, error) {
	result, ok := dictionary[searchWord]
	if !ok {
		return result, ErrSearchWordNotFound
	}
	return result, nil
}

func (dic Dictionary) Add(key, value string) error {
	_, err := dic.Search(key)

	switch err {
	case ErrSearchWordNotFound:
		dic[key] = value
	case nil:
		return ErrKeyAlreadyExist
	default:
		return err
	}
	return nil
}

func (dic Dictionary) Update(key, value string) error {
	_, err := dic.Search(key)
	switch err {
	case nil:
		dic[key] = value
	case ErrSearchWordNotFound:
		return ErrCannotUpdateNonExistingKey
	default:
		return err
	}
	return nil
}

func (dic Dictionary) Delete(key string) {
	delete(dic, key)
}
