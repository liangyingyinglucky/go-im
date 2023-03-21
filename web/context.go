package web

import (
	"net/http"
	"fmt"
)

//上下文参数传递
type Context struct {
	W http.ResponseWriter
	R *http.Request
	PathParams map[string]string
}

func newContext() *Context {
	fmt.Println("create new context")
	return &Context{
	}
}


func (c *Context) Reset(w http.ResponseWriter, r *http.Request) {
	c.W = w
	c.R = r
	c.PathParams = make(map[string]string, 1)
}