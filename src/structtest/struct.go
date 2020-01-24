package structtest

import (
	"fmt"
)

type Monitor struct {
	x int
	y int
	w int
	h int
}

func (m *Monitor) GetMonitorInfo() {
	fmt.Println("x :", m.x, ", y :", m.y, ", width :", m.w, ", height :", m.h)
}


