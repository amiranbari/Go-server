package config

import "html/template"

type TemplateCache map[string]*template.Template

type AppConfig struct {
	UseCache      bool
	TemplateCache TemplateCache
}
