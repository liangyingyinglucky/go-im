package web

import (
	"strings"
)

const (
	// 根节点，只有根用这个
	nodeTypeRoot = iota
	// *
	nodeTypeAny
	// 路径参数
	nodeTypeParam
	// 静态，即完全匹配
	nodeTypeStatic
)

const any = "*"

// matchFunc 承担两个职责，一个是判断是否匹配，一个是在匹配之后将必要的数据写入到 Context
//所谓必要的数据，这里基本上是指路径参数
type matchFunc func(path string, c *Context) bool

type node struct {
	children []*node
	// 如果这是叶子节点，
	// 那么匹配上之后就可以调用该方法
	handler   handlerFunc
	matchFunc matchFunc
	pattern string
	nodeType int
}

//创建根节点
func newRootNode(method string) *node {
	return &node{
		children: make([]*node, 0, 2),
		matchFunc: func( p string, c *Context) bool {
			return false//根节点不调用这个方法
		},
		nodeType: nodeTypeRoot,
		pattern:  method,
	}
}

//创建路由匹配节点
func newNode(path string) *node {
	if path == "*"{
		return newAnyNode()//通配符匹配
	}
	if strings.HasPrefix(path, ":") {//一般路径匹配
		return newParamNode(path)
	}
	return newStaticNode(path)//完全匹配
}

// 通配符 * 节点
func newAnyNode() *node {
	return &node{
		// 因为我们不允许 * 后面还有节点，所以这里可以不用初始化
		//children: make([]*node, 0, 2),
		matchFunc: func(p string, c *Context) bool {
			return true
		},
		nodeType: nodeTypeAny,
		pattern:  any,
	}
}

// 路径参数节点
func newParamNode(path string) *node {
	paramName := path[1:]
	return &node{
		children: make([]*node, 0, 2),
		matchFunc: func(p string, c *Context) bool {
			if c != nil {
				//路由参数传递
				c.PathParams[paramName] = p
			}
			// 如果自身是一个参数路由，
			// 然后又来一个通配符，我们认为是不匹配的
			return p != any
		},
		nodeType: nodeTypeParam,
		pattern:  path,
	}
}


//完全匹配节点
func newStaticNode(path string) *node {
	return &node{
		children: make([]*node, 0, 2),
		matchFunc: func(p string, c *Context) bool {
			return path == p && p != "*"
		},
		nodeType: nodeTypeStatic,
		pattern:  path,
	}
}


