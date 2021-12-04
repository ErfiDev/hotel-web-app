package utils

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/erfidev/hotel-web-app/config"
	"github.com/erfidev/hotel-web-app/models"
	"github.com/justinas/nosurf"
)

var funcMap = template.FuncMap{}

var appConfig *config.AppConfig

func GetAppConfig(a *config.AppConfig) {
	appConfig = a
}

func AddDefaultData(tmpData *models.TmpData, req *http.Request) *models.TmpData {
	tmpData.Error = appConfig.Session.PopString(req.Context(), "error")
	tmpData.Flash = appConfig.Session.PopString(req.Context(), "flash")
	tmpData.Warning = appConfig.Session.PopString(req.Context(), "warning")
	tmpData.CSRF = nosurf.Token(req)
	if appConfig.Session.Exists(req.Context(), "user_id") {
		tmpData.Auth = 1
	}
	return tmpData
}

func RenderTemplate(w http.ResponseWriter, req *http.Request, tmpl string, data *models.TmpData) {
	var tmpCache map[string]*template.Template

	if appConfig.Development {
		tmpCache, _ = CreateTemplateCache()
	} else {
		tmpCache = appConfig.TemplatesCache
	}

	find, ok := tmpCache[tmpl]
	if !ok {
		log.Fatal("can't find template")
	}

	data = AddDefaultData(data, req)

	find.Execute(w, data)
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	caches := map[string]*template.Template{}

	pages, err := filepath.Glob("./views/*.page.gohtml")
	// [$../views/about.page.gohtml  &...]
	if err != nil {
		return caches, err
	}

	// [$../views/about.page.gohtml]
	for _, page := range pages {
		name := filepath.Base(page)
		// [about.page.gohtml]

		tmp, errNewTmp := template.New(name).Funcs(funcMap).ParseFiles(page)
		if errNewTmp != nil {
			return caches, errNewTmp
		}

		findLayouts, _ := filepath.Glob("./views/*.layout.gohtml")
		if len(findLayouts) > 0 {
			if strings.Contains(name, "admin") {
				tmp.ParseGlob("./views/admin.layout.gohtml")
			} else {
				tmp.ParseGlob("./views/base.layout.gohtml")
			}
		}

		caches[name] = tmp
	}
	return caches, nil
}
