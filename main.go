package main

import(
	 _ "fmt"
	 "gutils"
)

func main() {
	tt := gutils.NewTree(".")

	tn := gutils.TreeRoot(tt)
	_ = tn.AddChild("t1")
	tn2 := tn.AddChild("t2")
	tn2.AddChild("ts1")
	tn2.AddChild("ts2")

	gutils.PrintTree(tt)
}
