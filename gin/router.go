package gin

import (
	"log"
	"net/http"
	"strings"
)

// roots key eg, roots['GET'] roots['POST']
// handlers key eg, handlers['GET-/p/:lang/doc'], handlers['POST-/p/book']
type Router struct {
	roots    map[string]*node       // key is method (GET, POST ...). Each method is corresponding to one root
	handlers map[string]HandlerFunc // key is method-pattern
}

func newRouter() *Router {
	return &Router{
		handlers: make(map[string]HandlerFunc),
		roots:    make(map[string]*node),
	}
}

func parsePattern(pattern string) []string {
	patternArray := strings.Split(pattern, "/")

	parts := make([]string, 0)

	for _, item := range patternArray {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' { // once *, then stopping checking more parts (Only * is allowed)
				break
			}
		}
	}

	return parts
}

func (r *Router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern // e.g. GET-/hello

	parts := parsePattern(pattern)

	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{} // create new root node if non-existed method
	}

	// add new trie for new method, and insert pattern into trie
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

// pattern /p/:age/doc, input /p/20/doc, then result is {age: "20"}
// pattern /p/*heightpath, input /p/170/weight, then result is {height: "170/weight"}
func (r *Router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	node := root.search(searchParts, 0)
	if node != nil {
		parts := parsePattern(node.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index] // {age: "20"}
			}

			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break // only assume one * in pattern
			}
		}
		return node, params
	}

	return nil, nil
}

func (r *Router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)

	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
