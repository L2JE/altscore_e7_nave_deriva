package common

type Line struct {
	slope float64
	y1    float64
	x1    float64
}

func NewLine(x1 float64, y1 float64, x2 float64, y2 float64) *Line {
	return &Line{
		slope: (y2 - y1) / (x2 - x1),
		x1:    x1,
		y1:    y1,
	}
}

func (l *Line) GetXCoord(y float64) float64 {
	return ((y - l.y1) / l.slope) + l.x1
}
