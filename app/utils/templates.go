package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// LoaderTemplates insere os templates htmls na variavel templates
func LoaderTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

// ExecutorTemplate renderiza uma pagina HTML na tela
func ExecutorTemplate(w http.ResponseWriter, template string, data interface{}) {
	templates.ExecuteTemplate(w, template, data)
}
