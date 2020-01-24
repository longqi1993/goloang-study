package runetest

import (
	"fmt"
	"unicode/utf8"
)

func printInfo(w string) {
	//sw := w[:2]
	//fmt.Println(sw)
	//sw[1] = '0'
	fmt.Printf("world :%s , len size : %d, rune size: %d\n",
		w, len(w), utf8.RuneCountInString(w))
	var wb []byte
	for i := 0; i < len(w); i++ {
		wb = append(wb, w[i])
	}
	fmt.Printf("world byte:")
	fmt.Println(wb)

	var wr []rune
	for _, r := range w {
		wr = append(wr, r)
	}
	fmt.Printf("world runes: ")
	fmt.Println(wr)
}

func RuneTest() {
	world1 := "helloword"
	world2 := "你好世界"
	_ = world2

	printInfo(world1)
	printInfo(world2)
}
