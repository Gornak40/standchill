package engine

import (
	"bufio"
	"log/slog"
	"net/http"
	"os/exec"
	"strings"

	"github.com/Gornak40/standchill/internal/utils"
)

const (
	loginsForm    = "logins"
	standingsTmpl = "standings.html"

	secretLink = "https://algocourses.ru/standings/d2024_fall_e26d813e808378e66bc0a2c1?logins="
)

func (e *Engine) StandingsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logins := r.FormValue(loginsForm)
		lr := bufio.NewScanner(strings.NewReader(logins))
		good := make(map[string]struct{})
		for lr.Scan() {
			good[lr.Text()] = struct{}{}
		}
		cmd := exec.Command("shoga", "-i", "55000")
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			utils.Report(w, err)
			return
		}
		if err := cmd.Start(); err != nil {
			utils.Report(w, err)
			return
		}
		csvr := bufio.NewScanner(stdout)
		for csvr.Scan() {
			slog.Info(csvr.Text())
		}
		if err := cmd.Wait(); err != nil {
			utils.Report(w, err)
			return
		}

		data := map[string]string{
			"Link": strings.TrimRight(secretLink, ","),
		}
		if err := e.tmpl.ExecuteTemplate(w, standingsTmpl, data); err != nil {
			utils.Report(w, err)
			return
		}
	}
}
