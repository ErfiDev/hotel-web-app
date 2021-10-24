package utils

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var funcMap = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter , tmpl string , data interface{}) {
	caches , err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	findTmp , isOk := caches[tmpl]
	if !isOk{
		log.Fatal("not found template")
	}

	findTmp.Execute(w , data)
}

func CreateTemplateCache() (map[string]*template.Template , error) {
	caches := map[string]*template.Template{}

	pages , err := filepath.Glob("./views/*.page.gohtml")
	// [$../views/about.page.gohtml  &...]
	if err != nil {
		return caches , err
	}

	// [$../views/about.page.gohtml]
	for _ , page := range pages {
		name := filepath.Base(page)
		// [about.page.gohtml]

		tmp , err := template.New(name).Funcs(funcMap).ParseFiles(page)
		if err !=  nil {
			return caches , err
		}

		findLayout , _ := filepath.Glob("./views/*.layout.gohtml")
		if len(findLayout) > 0 {
			tmp.ParseGlob("./views/*.layout.gohtml")
		}

		caches[name] = tmp
	}
	return caches , nil
}