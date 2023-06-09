package gee

import (
	"log"
	"net/http"
)

type HandlerFunc func(ctx *Context)

type Engine struct {
	router *router
}

func (engine *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//TODO implement me
	c := newContext(writer, request)
	engine.router.handle(c)
}

func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route:%s-%s\n", method, pattern)
	engine.router.addRouter(method, pattern, handler)
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
