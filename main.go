package main

import(
	 _ "fmt"
	 "gutils"
)

func main() {
	tt := gutils.NewTree(".")

	tn := tt.Root()
	tn.AddChild("t1")
	tn2 := tn.AddChild("t2")
	ttn1 := tn2.AddChild("ts1")
	ttn1.AddChild("tts1")
	ttn1.AddChild("tts2")
	ttn1.AddChild("tts3")
	tn2.AddChild("ts2")

	tt.PrintTree()
}
