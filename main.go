package main

import(
	_ "fmt"
	"io/ioutil"
	"strings"

	"gutils"
)

const (
	gsPath = "/usr/share/glib-2.0/schemas/"
)


func main() {
	gt := gutils.NewTree(" ")
	infos, _ := ioutil.ReadDir(gsPath)
	for _, v := range infos {
		addGsPath(v.Name(), gt)
	}

	gt.PrintTree()
}

func addGsPath(p string, t *gutils.Tree) {
	ss := strings.Split(p, ".")
	tn := t.Root()
	for idx := 0; idx < len(ss) - 2; idx++ {
		tn = tn.AddChild(ss[idx])
	}
}
