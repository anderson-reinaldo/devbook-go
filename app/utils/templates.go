package utils

import (
	"html/template"
	"net/http"
)

// LoaderTemplates mantida para compatibilidade — não faz mais nada
func LoaderTemplates() {}

// ExecutorTemplate renderiza uma pagina HTML na tela (recarrega templates a cada request)
func ExecutorTemplate(w http.ResponseWriter, tmplName string, data interface{}) {
	tmpl := template.Must(template.ParseGlob("views/*.html"))
	if err := tmpl.ExecuteTemplate(w, tmplName, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
