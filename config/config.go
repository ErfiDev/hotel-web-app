package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
)

type AppConfig struct {
	Development bool
	TemplatesCache map[string]*template.Template
	Session *scs.SessionManager
}
