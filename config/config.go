package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
	"log"
)

type AppConfig struct {
	Development bool
	TemplatesCache map[string]*template.Template
	Session *scs.SessionManager
	ErrorLog *log.Logger
	InfoLog *log.Logger
}
