package config

import "html/template"

type AppConfig struct {
	TemplatesCache map[string]*template.Template
}
