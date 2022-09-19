package main

import "strings"

type HandlerBasedTree struct{
	root *Node
}

type Node struct {
	path string
	children []*Node
	handler handlerFunc
}

func (h *HandlerBasedTree) ServerHTTP(c *Context) {
	panic("implement me")
}

func (h *HandlerBasedTree) Route(method string, pattern string, handleFun handlerFunc) {
	patt := strings.Trim(pattern,"/")	//处理路径的前后斜杆， /user/friends  /user/friends/  user/friends/
	paths := strings.Split(patt, "/")

	cur := h.root
	for index, path := range paths{
		mathchild, ok := h.findMathChild(cur, path)
		if ok{
			cur = mathchild
		} else {
			h.createSubTree(cur, paths[index:], handleFun)
		}
	}
}

func (h *HandlerBasedTree) findMathChild(root *Node, path string) (*Node, bool){
	for _, child := range root.children{
		if child.path == path{
			return child, true
		}
	}
	return nil, false
}

//
func (h *HandlerBasedTree) createSubTree(root *Node, paths []string, handlerFun handlerFunc) {
	cur := root
	for _, path := range paths{
		nn := newNode(path)
		cur.children = append(cur.children, nn)
		cur = nn
	}
	cur.handler = handlerFun
}

func newNode(path string) *Node {
	return &Node{
		path: path,
		children: make([]*Node, 0, 4),
	}
}