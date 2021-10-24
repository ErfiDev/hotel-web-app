package utils

import (
	"github.com/erfidev/hotel-web-app/config"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var funcMap = template.FuncMap{}

var appConfig *config.AppConfig

func GetAppConfig(a *config.AppConfig) {
	appConfig = a
}

func RenderTemplate(w http.ResponseWriter , tmpl string , data interface{}) {
	if appConfig.Development {
		tmpCache , _ := CreateTemplateCache()
		tmp , ok := tmpCache[tmpl]
		if !ok {
			log.Fatal("we can't find the template")
		}

		tmp.Execute(w , data)
	} else {
		caches := appConfig.TemplatesCache

		findTmp , isOk := caches[tmpl]
		if !isOk{
			log.Fatal("not found template")
		}

		findTmp.Execute(w , data)
	}
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

		tmp , errNewTmp := template.New(name).Funcs(funcMap).ParseFiles(page)
		if errNewTmp !=  nil {
			return caches , errNewTmp
		}

		findLayout , _ := filepath.Glob("./views/*.layout.gohtml")
		if len(findLayout) > 0 {
			tmp.ParseGlob("./views/*.layout.gohtml")
		}

		caches[name] = tmp
	}
	return caches , nil
}