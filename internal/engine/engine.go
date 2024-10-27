package engine

import "text/template"

const (
	tmplGlob = "templates/*.html"
)

type Engine struct {
	tmpl *template.Template
}

func New() *Engine {
	return &Engine{
		tmpl: template.Must(template.ParseGlob(tmplGlob)),
	}
}
