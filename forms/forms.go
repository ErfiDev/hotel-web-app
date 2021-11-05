package forms

import (
	"net/http"
	"net/url"
)

type Form struct {
	url.Values
	Errors errors
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

func New(data url.Values) *Form {
	return &Form {
		data,
		errors{},
	}
}

func (f *Form) Has(field string , req *http.Request) bool {
	check := req.Form.Get(field)
	if check == "" {
		f.Errors.Add(field , "this field can't be blank")
		return false
	}

	return true
}