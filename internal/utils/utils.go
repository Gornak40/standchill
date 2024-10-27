package utils

import (
	"log/slog"
	"os"
)

func ErrAttr(err error) slog.Attr {
	return slog.String("error", err.Error())
}

func Fatal(msg string, args ...any) {
	slog.Error(msg, args...)
	os.Exit(1)
}
