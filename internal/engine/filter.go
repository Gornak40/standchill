package engine

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Gornak40/standchill/internal/utils"
)

const (
	filterTmpl = "filter.html"
)

func (e *Engine) FilterHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logins := r.FormValue(loginsForm)
		lines := strings.Split(logins, "\n")
		filter := make([]string, 0, len(lines))
		for _, lg := range lines {
			filter = append(filter, fmt.Sprintf("login == '%s'", lg))
		}
		data := map[string]string{
			"Filter": strings.Join(filter, "||"),
		}
		if err := e.tmpl.ExecuteTemplate(w, filterTmpl, data); err != nil {
			utils.Report(w, err)
			return
		}
	}
}
