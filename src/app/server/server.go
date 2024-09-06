package server

import (
	endpoints "altscore_e7_nave_deriva/app/model/ship"
	"strconv"

	"log/slog"
	"net/http"
	"os"
)

type AdriftShipServer struct {
	Port                 string
	RespondWHTMLFragment bool
	shipService          *endpoints.ShipService
}

func (server *AdriftShipServer) Start() {
	server.shipService = endpoints.InitShipService(!server.RespondWHTMLFragment)
	SetupRoutes(server.shipService)

	slog.Info("Starting AdriftShipServer at: " + server.Port)
	http.ListenAndServe(":"+server.Port, nil)
}

func NewAdriftShipServer() *AdriftShipServer {
	config := GenerateValidConfigFromENV()

	return &AdriftShipServer{
		Port:                 config.port,
		RespondWHTMLFragment: config.useHtmlFragment,
	}
}

type Config struct {
	port            string
	useHtmlFragment bool
}

func GenerateValidConfigFromENV() *Config {
	const APP_PORT_ENV = "APP_PORT"
	const APP_HTML_FRAMENTS = "APP_HTML_FRAMENTS"

	port := os.Getenv(APP_PORT_ENV)

	if port == "" {
		startupPanic(APP_PORT_ENV + " environment variable is not set")
	}

	useFragments, err := strconv.ParseBool(os.Getenv(APP_HTML_FRAMENTS))

	if err != nil {
		startupPanic(APP_HTML_FRAMENTS + " environment variable is not set or is an invalid value")
	}

	return &Config{port, useFragments}
}

func startupPanic(msg string) {
	slog.Error("Application cannot start: " + msg)
	panic("")
}
