package engine

import (
	"bufio"
	"encoding/csv"
	"errors"
	"io"
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

func (e *Engine) getLink(filter map[string]struct{}) (string, error) {
	cmd := exec.Command("shoga", "-i", "55000", "-m", "usr")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	if err := cmd.Start(); err != nil {
		return "", err
	}

	csvr := csv.NewReader(stdout)
	csvr.Comma = ';'
	link := secretLink
	for {
		rec, err := csvr.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return "", err
		}
		if len(rec) < 3 {
			slog.Warn("bad csv line", slog.Any("record", rec))
			continue
		}
		login, name := rec[1], rec[2]
		if _, ok := filter[login]; ok {
			link += name + ","
		}
	}
	if err := cmd.Wait(); err != nil {
		return "", err
	}
	return strings.TrimRight(link, ","), nil
}

func (e *Engine) StandingsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logins := r.FormValue(loginsForm)
		lr := bufio.NewScanner(strings.NewReader(logins))
		filter := make(map[string]struct{})
		for lr.Scan() {
			filter[lr.Text()] = struct{}{}
		}
		link, err := e.getLink(filter)
		if err != nil {
			utils.Report(w, err)
			return
		}
		data := map[string]string{
			"Link": link,
		}
		if err := e.tmpl.ExecuteTemplate(w, standingsTmpl, data); err != nil {
			utils.Report(w, err)
			return
		}
	}
}
