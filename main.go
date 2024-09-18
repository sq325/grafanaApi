package main

import (
	"github.com/sq325/grafanaApi/cmd"
)

func main() {
	cmd.Execute()
}

// func initLogger() {
// 	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
// 		Level:     slog.LevelInfo,
// 		AddSource: true,
// 	})

// 	logger := slog.New(handler)
// 	slog.SetDefault(logger)
// }
