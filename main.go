package main

import(
	"fmt"
	"io/ioutil"
	"gutils"
	"os"
	"strconv"
)

var (
	_all=false
	_dirOnly=false
	_maxDep=-1
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("need dir args")
	}

	dir := args[len(args) -1]
	file, err := os.Open(dir)
	if err != nil {
		fmt.Printf("read file %s faild! error :", dir, err.Error())
		return
	}

	for idx := 1; idx < len(args) -1; idx++ {
		if args[idx][0] == '-' && len(args[idx]) > 1 {
			for ci := 1; ci < len(args[idx]); ci++ {
				switch(args[idx][ci]) {
				case 'a','A':
					_all = true
				case 'd','D':
					_dirOnly = true
				default:
					fmt.Printf("can not resolve args %1 !\n", args[idx][ci])
				}
			}
		} else {
			dp, err := strconv.Atoi(args[idx])
			if err != nil {
				fmt.Println("the dir args mast at last")
				return
			}
			_maxDep = dp
		}
	}

	ft := gutils.NewTree(file.Name())
	addFile(dir, ft.Root())
	ft.PrintTree()
}

func addFile(dir string, tn *gutils.TreeNode) {
	infos, _ := ioutil.ReadDir(dir)
	//fmt.Println("find dir :", dir, ", get file size :", len(infos))

	for _, cf := range infos {
		if cf.Name()[0] == '.' && !_all {
			continue
		}
		//fmt.Println("add Child name :", cf.Name())
		if (!_dirOnly && !cf.IsDir()) {
			tn.AddChild(cf.Name())
		} else if (cf.IsDir()) {
			ctn := tn.AddChild(cf.Name())
			addFile(dir + "/" + cf.Name(), ctn)
		}
	}
}
