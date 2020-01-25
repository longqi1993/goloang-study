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

func NewTree(s string) *Tree {
	tn := &TreeNode{_p:nil, _c:nil, _d:[]rune(s)}
	return &Tree{_r:tn}
}

func (t *Tree) PrintTree() {
	t._r.PrintNode()
}

func (tn *TreeNode) PrintNode() {
	fmt.Printf(string(tn._d))
}
