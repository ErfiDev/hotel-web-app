package forms

import (
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
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

func (f *Form) Has(field string, req *http.Request) bool {
	value := req.Form.Get(field)
	if value == "" {
		f.Errors.Add(field , "this field can't be blank")
		return false
	} else {
		if field == "first_name" || field == "last_name" {
			isOk := CheckNameWithRegex(value)
			if !isOk {
				f.Errors.Add(field , "just can use a-z letters")
				return false
			} else {
				return true
			}
		}

		if field == "email" {
			isOk := CheckEmail(value)
			if !isOk {
				f.Errors.Add(field , "field must be a valid email")
				return false
			} else {
				return true
			}
		}

		if field == "phone" {
			isOk := CheckPhone(value)
			if !isOk {
				f.Errors.Add(field , "field must be a valid phone")
				return false
			} else {
				return true
			}
		}

		if field == "start-date" {
			isOk := CheckStartTime(value)
			if !isOk {
				f.Errors.Add("start-date" , "please define the valid date")
				return false
			} else {
				return true
			}
		}

		if field == "ending-date" {
			isOk := CheckEndTime(value)
			if !isOk {
				f.Errors.Add("ending-date" , "please define the valid date")
				return false
			} else {
				return true
			}
		}

		return true
	}
}

func CheckNameWithRegex(data string) bool {
	reg := regexp.MustCompile("^[a-zA-Z]+(([',. -][a-zA-Z ])?[a-zA-Z]*)*$")

	return reg.Match([]byte(data))
}

func CheckEmail(email string) bool {
	reg := regexp.MustCompile("^((\"[\\w-\\s]+\")|([\\w-]+(?:\\.[\\w-]+)*)|(\"[\\w-\\s]+\")([\\w-]+(?:\\.[\\w-]+)*))(@((?:[\\w-]+\\.)*\\w[\\w-]{0,66})\\.([a-z]{2,6}(?:\\.[a-z]{2})?)$)|(@\\[?((25[0-5]\\.|2[0-4][0-9]\\.|1[0-9]{2}\\.|[0-9]{1,2}\\.))((25[0-5]|2[0-4][0-9]|1[0-9]{2}|[0-9]{1,2})\\.){2}(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[0-9]{1,2})\\]?$)")

	return reg.Match([]byte(email))
}

func CheckPhone(phone string) bool {
	reg := regexp.MustCompile("^[0-9]*$")

	return reg.Match([]byte(phone))
}

func CheckStartTime(start string) bool {
	splitStart := TimeSpliter(start)

	toDateStart := time.Date(splitStart[0] , time.Month(splitStart[2]), splitStart[1] , 00 , 01 , 01 , 80 , &time.Location{})
	now := time.Now()

	if !now.After(toDateStart){
		return true
	} else {
		return false
	}
}

func CheckEndTime(start string) bool {
	splitEnd := TimeSpliter(start)

	toDateEnd := time.Date(splitEnd[0] , time.Month(splitEnd[2]), splitEnd[1] , 00 , 01 , 01 , 80 , &time.Location{})
	now := time.Now()

	if !now.After(toDateEnd){
		return true
	} else {
		return false
	}
}

func TimeSpliter(time string) []int {
	split := strings.Split(time, "")

	yy , _ := strconv.Atoi(strings.Join(split[:4] , ""))
	mm , _ := strconv.Atoi(strings.Join(split[5:7] , ""))
	dd , _ := strconv.Atoi(strings.Join(split[8:10] , ""))

	return []int{yy , dd , mm}
}