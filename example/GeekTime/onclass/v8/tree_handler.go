package main

import (
	"net/http"
	"strings"
)

type HandlerBasedTree struct {
	root *Node
}

type Node struct {
	path     string
	children []*Node
	handler  handlerFunc
}

func (h *HandlerBasedTree) ServerHTTP(c *Context) {
	handler, found := h.findRouter(c.R.URL.Path)
	if !found {
		c.W.WriteHeader(http.StatusNotFound)
		_, _ = c.W.Write([]byte("Not Found"))
		return
	}
	handler(c)
}

func (h *HandlerBasedTree) findRouter(path string) (handlerFunc, bool) {
	// 去除头尾可能有的/，然后按照/切割成段
	paths := strings.Split(strings.Trim(path, "/"), "/")
	cur := h.root
	for _, p := range paths {
		// 从子节点里边找一个匹配到了当前 path 的节点
		matchChild, found := h.findMathChild(cur, p)
		if !found {
			return nil, false
		}
		cur = matchChild
	}
	// 到这里，应该是找完了
	if cur.handler == nil {
		// 到达这里是因为这种场景
		// 比如说你注册了 /user/profile
		// 然后你访问 /user
		return nil, false
	}
	return cur.handler, true
}

func (h *HandlerBasedTree) Route(method string, pattern string, handleFun handlerFunc) {
	patt := strings.Trim(pattern, "/") //处理路径的前后斜杆， 例如将：/user/friends  /user/friends/  user/friends/  转换成：user/friends
	paths := strings.Split(patt, "/")  //然后将user/friends切分为[user, friends]

	cur := h.root
	for index, path := range paths {
		mathchild, ok := h.findMathChild(cur, path)
		if ok {
			cur = mathchild
		} else { //若为匹配到，创建路由树
			h.createSubTree(cur, paths[index:], handleFun)
		}
	}
}

func (h *HandlerBasedTree) findMathChild(root *Node, path string) (*Node, bool) {
	for _, child := range root.children {
		if child.path == path {
			return child, true
		}
	}
	return nil, false
}

//创建路由树
func (h *HandlerBasedTree) createSubTree(root *Node, paths []string, handlerFun handlerFunc) {
	cur := root
	for _, path := range paths {
		nn := newNode(path) //创建新节点
		cur.children = append(cur.children, nn)
		cur = nn
	}
	cur.handler = handlerFun
}

//创建新节点
func newNode(path string) *Node {
	return &Node{
		path:     path,
		children: make([]*Node, 0, 4),
	}
}
