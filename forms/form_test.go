package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForms(t *testing.T) {
	req := httptest.NewRequest("POST" , "/" , nil)

	form := New(url.Values{})

	form.Has("name" , req)

	if form.Valid() {
		t.Error("form testing failed!")
	}
}