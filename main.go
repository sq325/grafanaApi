package main

import (
	"log/slog"
	"os"

	"github.com/sq325/grafanaApi/cmd"
)

func main() {
	initLogger()
	cmd.Execute()
}

func initLogger() {
	handlerOpt := &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}
			if a.Key == slog.LevelKey {
				return slog.Attr{}
			}
			return a
		},
		AddSource: true,
	}
	handler := slog.NewTextHandler(os.Stdout, handlerOpt)
	logger := slog.New(handler)
	slog.SetDefault(logger)
}
