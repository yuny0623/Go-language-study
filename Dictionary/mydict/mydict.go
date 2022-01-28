package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string // 이건 struct가 아니다. 그냥 map[string]string에 대해서 alias를 만들어 줄 수 있는거다.

var ( // 이런 식으로 한번에 선언하는 것도 가능하다.
	errNotFound   = errors.New("Not Found")
	errCantUpdate = errors.New("Cant update non-existing word")
	errWordExists = errors.New("That word already exists")
)

// type에 method를 추가할 수 있다는 장점이 있다.

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word] // go의 굉장히 편리한 기능중 하나이다. word가 있는지 찾으면 그게 두개 값을 리턴을 하고 value와 존재여부를 리턴해준다.
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

// Update a word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) {
	delete(d, word) // 이건 기본으로 제공되는 함수이다.
}
