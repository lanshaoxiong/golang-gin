package gin

import (
	"net/http"
)

// Custom request handler
type HandlerFunc func(*Context)

// Uniform handler for all http requests
type Engine struct {
	router *Router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// start a http server
func (engine *Engine) Run(port string) (err error) {
	return http.ListenAndServe(port, engine)
}

// the defined function of handler:
// type Handler interface {
// 	ServeHTTP(ResponseWriter, *Request)
// }
func (engine *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	c := newContext(writer, request) // for each http request, we will trigger new context. Serve in parallel
	engine.router.handle(c)
}
