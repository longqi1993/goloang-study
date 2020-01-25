package gutils

import (
	"fmt"
	_ "unicode/utf8"
)

type TreeNode struct {
	_p *TreeNode
	_c []*TreeNode
	_d []rune
}

type Tree struct {
	_r *TreeNode
}

const (
	lline = "└── "
	bline = "├── "
	pline = "│   "
	eline = "    "
)

func (tn *TreeNode) findChild(d string) bool {
	if (tn == nil) {
		return false
	}

	for _, v := range tn._c {
		if string(v._d) == d {
			return true;
		}
	}

	return false
}

func (tn *TreeNode) printChild(ps []rune) {
	for idx, v := range tn._c {
		fmt.Printf(string(ps))
		var tps []rune
		if idx == len(tn._c) - 1 {
			fmt.Printf(lline)
			tps = append(ps, []rune(eline)...)
		} else {
			fmt.Printf(bline)
			tps = append(ps, []rune(pline)...)
		}
		fmt.Printf(string(v._d))
		fmt.Println()
		if len(v._c) > 1 {
			v.printChild(tps)
		}
	}
}

func NewTree(s string) *Tree {
	tn := &TreeNode{_p:nil, _c:nil, _d:[]rune(s)}
	return &Tree{_r:tn}
}

func (t *Tree) PrintTree() {
	t._r.PrintNode()
}

func (t *Tree) Root() *TreeNode {
	return t._r
}

func (tn *TreeNode) PrintNode() {
	fmt.Printf(string(tn._d))
	fmt.Println()
	tn.printChild([]rune{})
}

func (tn *TreeNode) AddChild (d string) *TreeNode {
	tnode := &TreeNode{_p:tn, _c:nil, _d:[]rune(d)}
	tn._c = append(tn._c, tnode)
	return tnode
}



