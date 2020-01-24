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

func NewMonitor() *Monitor {
	m := Monitor{}
	return &m
}

func (m *Monitor) GetMonitorInfo() {
	fmt.Println("x :", m.x, ", y :", m.y, ", width :", m.w, ", height :", m.h)
}


