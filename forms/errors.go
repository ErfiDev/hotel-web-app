package forms

type errors map[string][]string

func (e errors) Add(field , msg string){
	e[field] = append(e[field] , msg)
}

func (e errors) Get(field string) string {
	sliceValue := e[field]
	if len(sliceValue) == 0 {
		return ""
	}

	return sliceValue[0]
}
