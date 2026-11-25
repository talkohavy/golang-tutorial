package jsontags

import (
	"encoding/json"
	"testing"

	"go_tutorial/internal/testkit"
)

func TestBookType(t *testing.T) {
	tests := []struct {
		b    Book
		want string
	}{
		{Book{Title: "Go", Author: "Rob"}, `{"title":"Go","author":"Rob"}`},
		{Book{Title: "Go", Author: "Rob", Pages: 300}, `{"title":"Go","author":"Rob","pages":300}`},
	}
	for _, tt := range tests {
		data, err := json.Marshal(tt.b)
		if got := string(data); got != tt.want || err != nil {
			t.Errorf("json.Marshal(%+v) = %s, %v, want %s, nil", tt.b, got, err, tt.want)
		}
	}
}

func TestToJSON(t *testing.T) {
	tests := []struct {
		b    Book
		want string
	}{
		{Book{Title: "Go", Author: "Rob"}, `{"title":"Go","author":"Rob"}`},
		{Book{Title: "Dune", Author: "Frank", Pages: 412}, `{"title":"Dune","author":"Frank","pages":412}`},
	}
	for _, tt := range tests {
		toJSON := testkit.Method[func() (string, error)](t, &tt.b, "ToJSON")
		if got, err := toJSON(); got != tt.want || err != nil {
			t.Errorf("%+v.ToJSON() = %s, %v, want %s, nil", tt.b, got, err, tt.want)
		}
	}
}

func TestFromJSON(t *testing.T) {
	var b Book
	fromJSON := testkit.Method[func(string) error](t, &b, "FromJSON")
	if err := fromJSON(`{"title":"Dune","author":"Frank","pages":412}`); err != nil {
		t.Fatalf("FromJSON(...) returned error %v, want nil", err)
	}
	switch want := (Book{Title: "Dune", Author: "Frank", Pages: 412}); b {
	case want:
	case Book{}:
		t.Error("after FromJSON(...) the book is still empty: did FromJSON decode into a copy?")
	default:
		t.Errorf("after FromJSON(...) the book is %+v, want %+v", b, want)
	}

	if err := fromJSON(`{not json`); err == nil {
		t.Error("FromJSON(`{not json`) returned no error")
	}
}
