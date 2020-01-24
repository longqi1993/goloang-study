package main

import(
	 _ "fmt"
	 "structtest"
)

func main(){
	monitor := structtest.Monitor{
		0,
		0,
		0,
		0,
	}

	monitor.GetMonitorInfo()
}
