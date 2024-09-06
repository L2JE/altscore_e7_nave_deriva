package main

import (
	"log/slog"
	"os"

	server "altscore_e7_nave_deriva/app/server"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	server := server.NewAdriftShipServer()
	server.Start()
}
