package utils

import (
	"log/slog"
	"net/http"
	"os"
)

func ErrAttr(err error) slog.Attr {
	return slog.String("error", err.Error())
}

func Fatal(msg string, args ...any) {
	slog.Error(msg, args...)
	os.Exit(1)
}

func Report(w http.ResponseWriter, err error) {
	slog.Error("internal server error", ErrAttr(err))
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
