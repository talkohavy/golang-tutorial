//go:build solutions

// SPOILER ALERT: reference solution. Try the exercise first!

package jsontags

import "encoding/json"

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages,omitempty"`
}

func (b Book) ToJSON() (string, error) {
	data, err := json.Marshal(b)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// FromJSON needs a pointer receiver to fill in the caller's Book, and passes
// that pointer on to json.Unmarshal for the same reason.
func (b *Book) FromJSON(s string) error {
	return json.Unmarshal([]byte(s), b)
}
