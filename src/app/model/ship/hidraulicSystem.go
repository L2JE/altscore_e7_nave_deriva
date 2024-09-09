package ship

import (
	geometry "altscore_e7_nave_deriva/utils/common"
)

type HidraulicSystem struct {
	saturatedLiquidLine *geometry.Line
	saturatedVaporLine  *geometry.Line
}

type PhaseChangeSpecificVolume struct {
	VolumeLiquid float64 `json:"specific_volume_liquid"`
	VolumeVapor  float64 `json:"specific_volume_vapor"`
}

type PhaseChangeLines struct {
	LiquidX1 float64
	LiquidY1 float64
	LiquidX2 float64
	LiquidY2 float64
	VaporX1  float64
	VaporY1  float64
	VaporX2  float64
	VaporY2  float64
}

func NewHidraulicSystem(phaseChangeDiagram *PhaseChangeLines) *HidraulicSystem {
	return &HidraulicSystem{
		saturatedLiquidLine: geometry.NewLine(
			phaseChangeDiagram.LiquidX1,
			phaseChangeDiagram.LiquidY1,
			phaseChangeDiagram.LiquidX2,
			phaseChangeDiagram.LiquidY2,
		),
		saturatedVaporLine: geometry.NewLine(
			phaseChangeDiagram.VaporX1,
			phaseChangeDiagram.VaporY1,
			phaseChangeDiagram.VaporX2,
			phaseChangeDiagram.VaporY2,
		),
	}
}

func (hs *HidraulicSystem) GetPhaseChangeValuesForPressure(pressure float64) *PhaseChangeSpecificVolume {
	return &PhaseChangeSpecificVolume{
		VolumeLiquid: hs.saturatedLiquidLine.GetXCoord(pressure),
		VolumeVapor:  hs.saturatedVaporLine.GetXCoord(pressure),
	}
}
