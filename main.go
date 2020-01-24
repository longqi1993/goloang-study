package main

import(
	 _ "fmt"
	 "structtest"
)

func main() {
	m := structtest.NewMonitor()

	m.GetMonitorInfo()
}
