package utils

import (
	"github.com/erfidev/hotel-web-app/config"
	"github.com/erfidev/hotel-web-app/models"
	"github.com/justinas/nosurf"
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

func AddDefaultData(tmpData *models.TmpData , req *http.Request) *models.TmpData {
	tmpData.CSRF = nosurf.Token(req)
	return tmpData
}

func RenderTemplate(w http.ResponseWriter , req *http.Request , tmpl string , data *models.TmpData) {
	var tmpCache map[string]*template.Template

	if appConfig.Development {
		tmpCache , _ = CreateTemplateCache()
	} else {
		tmpCache = appConfig.TemplatesCache
	}

	find , ok := tmpCache[tmpl]
	if !ok {
		log.Fatal("can't find template")
	}

	data = AddDefaultData(data , req)

	find.Execute(w , data)
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