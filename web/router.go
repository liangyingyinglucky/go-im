package web

import (
	"errors"
	"net/http"
	"sort"
	"strings"
	"fmt"
)

//错误信息
var ErrorInvalidRouterPattern = errors.New("invalid router pattern")
var ErrorInvalidMethod = errors.New("invalid method")

//区分请求的方法 post get分别所属的树
type HandlerBasedOnTree struct {
	forest map[string]*node
}

//请求类型
var supportMethods = [4]string {
	http.MethodGet,
	http.MethodPost,
	http.MethodPut,
	http.MethodDelete,
}


func NewHandlerBasedOnTree() Handler {
	forest := make(map[string]*node, len(supportMethods))
	for _, m :=range supportMethods {
		forest[m] = newRootNode(m)
	}
	return &HandlerBasedOnTree{
		forest: forest,
	}
}


//从树里面找节点 找到了就执行
func (h *HandlerBasedOnTree) ServeHTTP(c *Context) {
	handler, found := h.findRouter(c.R.Method, c.R.URL.Path, c)
	if !found {
		c.W.WriteHeader(http.StatusNotFound)
		_, _ = c.W.Write([]byte("Not Found"))
		return
	}
	handler(c)
}

//查找路由
func (h *HandlerBasedOnTree) findRouter(method, path string, c *Context) (handlerFunc, bool) {
	//格式化路径，看走的是什么树
	paths := strings.Split(strings.Trim(path, "/"), "/")
	//是否能匹配到某个树
	cur, ok := h.forest[method]
	if !ok {
		return nil, false
	}

	//节点层层找，直到找到最后一层
	for _, p := range paths {
		// 从子节点里边找一个匹配到了当前路由节点
		matchChild, found := h.findMatchChild(cur, p, c)
		if !found {
			return nil, false
		}
		cur = matchChild
	}
	//找不到
	if cur.handler == nil {
		return nil, false
	}
	return cur.handler, true
}

//匹配当前路由
func (h *HandlerBasedOnTree) findMatchChild(root *node,
	path string, c *Context) (*node, bool) {
	matchChild := make([]*node, 0, 2)
	for _, child := range root.children {
		//当前是否匹配
		if child.matchFunc(path, c) {
			matchChild = append(matchChild, child)
		}
	}
	if len(matchChild) == 0 {
		return nil, false
	}
	//多个匹配时，怎么确定使用哪种匹配
	//通过nodeType的前后关系决定决定优先级，然后升序排列
	sort.Slice(matchChild, func(i, j int) bool {
		return matchChild[i].nodeType < matchChild[j].nodeType
	})
	//将完全匹配放第一个返回
	return matchChild[len(matchChild) - 1], true
}

//往树中注册节点
func (h *HandlerBasedOnTree) Route(method string, pattern string,
	handlerFunc handlerFunc) error {
	//是否符合既定的路由规则，不符合规则的不让注册
	err := h.validatePattern(pattern)
	if err != nil {
		return err
	}
	// 将pattern按照URL的分隔符切割
	pattern = strings.Trim(pattern, "/")
	paths := strings.Split(pattern, "/")

	fmt.Println(paths)

	//当前走什么树，get  post 没有匹配则错误
	cur, ok := h.forest[method]
	if !ok {
		return ErrorInvalidMethod
	}
	for index, path := range paths {
	    //从当前节点去匹配
		matchChild, found := h.findMatchChild(cur, path, nil)
		//如果是先匹配*号节点，则可能导致后续节点匹配不上，所以要先判断一下
		if found && matchChild.nodeType != nodeTypeAny {
			cur = matchChild
		} else {
			//如果没有则注册节点
			h.createSubTree(cur, paths[index:], handlerFunc)
			return nil
		}
	}
	//没有循环就是只有一层路径的情况
	cur.handler = handlerFunc
	return nil
}

//自定义路由规则
func (h *HandlerBasedOnTree) validatePattern(pattern string) error {
	// 校验 *，如果存在，必须在最后一个，并且它前面必须是/
	// 即我们只接受 /* 的存在，abc*这种是非法
	pos := strings.Index(pattern, "*")
	// 找到了 *
	if pos > 0 {
		// 必须是最后一个
		if pos != len(pattern) - 1 {
			return ErrorInvalidRouterPattern
		}
		if pattern[pos-1] != '/' {
			return ErrorInvalidRouterPattern
		}
	}
	return nil
}


//创建子节点
func (h *HandlerBasedOnTree) createSubTree(root *node, paths []string, handlerFn handlerFunc) {
	cur := root//根节点，在一开始NewHandlerBasedOnTree的时候注册
	for _, path := range paths {
		nn := newNode(path)
		cur.children = append(cur.children, nn)
		cur = nn
	}
	cur.handler = handlerFn
}

