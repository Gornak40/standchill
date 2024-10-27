package engine

import (
	"net/http"

	"github.com/Gornak40/standchill/internal/utils"
)

const (
	indexTmpl = "index.html"
)

func (e *Engine) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := e.tmpl.ExecuteTemplate(w, indexTmpl, nil); err != nil {
			utils.Report(w, err)
			return
		}
	}
}
