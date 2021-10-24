package config

import "html/template"

type AppConfig struct {
	Development bool
	TemplatesCache map[string]*template.Template
}
