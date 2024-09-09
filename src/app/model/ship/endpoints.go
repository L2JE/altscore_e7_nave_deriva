package ship

import (
	mapUtils "altscore_e7_nave_deriva/utils/common"
	customJson "altscore_e7_nave_deriva/utils/middleware"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
)

func InitShipService(respondFullHTML bool) *ShipService {
	systems := map[string]string{
		"navigation":       "NAV-01",
		"communications":   "COM-02",
		"life_support":     "LIFE-03",
		"engines":          "ENG-04",
		"deflector_shield": "SHLD-05",
	}

	return &ShipService{
		systems,
		mapUtils.PickRandomKey(systems),
		respondFullHTML,
		NewHidraulicSystem(&PhaseChangeLines{
			LiquidX1: 0.00105,
			LiquidY1: 0.05,
			LiquidX2: 0.0035,
			LiquidY2: 10,
			VaporX1:  0.0035,
			VaporY1:  10,
			VaporX2:  30,
			VaporY2:  0.05,
		}),
	}
}

type ShipService struct {
	systems         map[string]string
	damagedSystem   string
	respondFullHTML bool
	hidraulicSystem *HidraulicSystem
}

type ShipStatus struct {
	DamagedSystem string `json:"damaged_system"`
}

func (ship *ShipService) GetShipStatus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&ShipStatus{ship.damagedSystem})
}

func (ship *ShipService) GetFailingSystem(w http.ResponseWriter, r *http.Request) {
	var textResponse string
	if ship.respondFullHTML {
		textResponse = `
<!DOCTYPE html>
<html>
<body>
<div class="anchor-point">%s</div>
</body>
</html>`
	} else {
		textResponse = `<div class="anchor-point">%s</div>`
	}

	textResponse = fmt.Sprintf(textResponse, ship.systems[ship.damagedSystem])

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err := fmt.Fprint(w, textResponse)

	if err != nil {
		w.WriteHeader(500)
	}
}

func (ship *ShipService) ImTeapotHealthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusTeapot)
}

func (ship *ShipService) GetPhaseChangeDiagramValues(w http.ResponseWriter, r *http.Request) {
	const PRESSURE_PARAM = "pressure"
	const JSON_FLOATS_PRECISSION = 5
	pressureStr := r.URL.Query().Get(PRESSURE_PARAM)
	pressure, err := strconv.ParseFloat(pressureStr, 64)

	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	customJson.EncodeRoundingFloats(w, ship.hidraulicSystem.GetPhaseChangeValuesForPressure(pressure), JSON_FLOATS_PRECISSION)
}
