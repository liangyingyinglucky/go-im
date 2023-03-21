package web

import (
	//"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

//Handler 自己的，为了抽象路由
type Handler interface {
	ServeHTTP(c *Context)
	Routable
}

type handlerFunc func(c *Context)

//最外层路由抽象
type Routable interface {
	//路由命中设置
	Route(method string,pattern string,handlerFunc handlerFunc) error
}

//server抽象
type Server interface {
	Routable
	Start(address string) error
}


type sdkHttpServer struct {
	Name string 
	handler Handler
	root Filter
	ctxPool sync.Pool
}

//注册路由
func (s *sdkHttpServer) Route(method string,pattern string,handler handlerFunc) error{
	fmt.Printf("router: %d \n", method)
	fmt.Printf("pattent: %d \n", pattern)
	return s.handler.Route(method,pattern,handler)
}



//监听
func (s *sdkHttpServer) Start(address string) error {
	return http.ListenAndServe(address,s)
}

//sdkHttpServer中的http.Handler的ServeHTTP
func (s *sdkHttpServer) ServeHTTP(writer http.ResponseWriter, request *http.Request)  {
	c := s.ctxPool.Get().(*Context)
	defer func() {
		s.ctxPool.Put(c)
	}()
	c.Reset(writer, request)
	s.root(c)
}

type FilterBuilder func(next Filter) Filter

type Filter func(c *Context)


//执行时间
func MetricFilterBuilder(next Filter) Filter {
	return func(c *Context) {
		// 执行前的时间
		startTime := time.Now().UnixNano()
		next(c)
		// 执行后的时间
		endTime := time.Now().UnixNano()
		fmt.Printf("run time: %d \n", endTime-startTime)
	}
}



func NewSdkHttpServer(name string, builders ...FilterBuilder) Server {
	handler := NewHandlerBasedOnTree()
	var root Filter = handler.ServeHTTP
	// 从后往前把filter串起来
	for i := len(builders) - 1; i >= 0; i-- {
		b := builders[i]
		root = b(root)
	}
	res := &sdkHttpServer{
		Name: name,
		handler: handler,
		root: root,
		ctxPool: sync.Pool{New: func() interface {}{
			return newContext()
		}},
	}
	return res
}



















