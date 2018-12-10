package api

import "testing"

func TestBookToJSON(t *testing.T) {
	book := Book{Title: "My Journey", Author: "Filip K", ISBN: "1234567"}
	data := book.ToJSON()
	if string(data) != `{"title":"My Journey","author":"Filip K","isbn":"1234567"}` {
		t.Errorf(`JSON dont match. Got: %s, expected {"title":"My Journey","author":"Filip K","isbn":"1234567"}`, string(data))
	}
}

func TestBookFromJSON(t *testing.T) {
	data := []byte(`{"title":"My Journey","author":"Filip K","isbn":"1234567"}`)
	book := BookFromJSON(data)
	if book.Title != "My Journey" {
		t.Errorf("Expected same title but got different %s", book.Title)
	}
	if book.Author != "Filip K" {
		t.Errorf("Expected same author but got different %s", book.Author)
	}
	if book.ISBN != "1234567" {
		t.Errorf("Expected same isbn but got different %s", book.ISBN)
	}
}
