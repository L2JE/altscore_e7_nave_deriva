package server

import (
	credit "altscore_e7_nave_deriva/app/banking"
	endpoints "altscore_e7_nave_deriva/app/model/ship"
	middleware "altscore_e7_nave_deriva/utils/middleware"
	"net/http"
)

func SetupRoutes(shipService *endpoints.ShipService) {
	mw := middleware.CreateMiddleware(middleware.LogIncomingRequests)

	http.Handle("GET /creditscore", mw.Apply(credit.GetCreditScore))
	http.Handle("GET /status", mw.Apply(shipService.GetShipStatus))
	http.Handle("GET /repair-bay", mw.Apply(shipService.GetFailingSystem))
	http.Handle("POST /teapot", mw.Apply(shipService.ImTeapotHealthcheck))
}
