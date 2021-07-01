package gin

import (
	"fmt"
	"net/http"
)

// Custom request handler
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Uniform handler for all http requests
type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern // e.g. GET-/hello
	engine.router[key] = handler
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

// the defined function of handler
// type Handler interface {
// 	ServeHTTP(ResponseWriter, *Request)
// }
func (engine *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	key := request.Method + "-" + request.URL.Path

	if handler, ok := engine.router[key]; ok { // ok is bool to indicate whether key exists https://blog.golang.org/maps
		handler(writer, request)
	} else {
		fmt.Fprintf(writer, "404 NOT FOUND: %s\n", request.URL)
	}

}
