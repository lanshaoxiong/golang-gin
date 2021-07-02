package gin

import "strings"

type node struct {
	pattern  string // e.g. /p/:age
	part     string // e.g. :age
	children []*node
	isWild   bool // it's true when it includes : or *
}

// match first child which matches part
func (n *node) matchChild(part string) *node {
	for _, child := range n.children { // ignore index value _
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// find all children which matches part
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0) // create a *node array. the size is 0

	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}

	return nodes
}

// insert new pattern(parts) into trie
// corresponding to addRoute
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height] // for root node, the height is 0
	child := n.matchChild(part)

	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'} // check whether this part contains * or :
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

// check whether trie has such parts
// corresponding to getRoute
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") { // check whether current node is * node
		if n.pattern == "" { // path is match, but we don't have handler for this path
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
