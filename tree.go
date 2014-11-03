package radix

import (
	"fmt"
	"io"
	"os"
)

type node struct {
	path     string
	children []*node
	value    interface{}
}

func New() *node {
	return &node{}
}

func (n *node) Insert(path string, value interface{}) {
	var ok bool
	for _, child := range n.children {
		ok = insert(path, value, child)
		if ok {
			break
		}
	}

	if !ok {
		sibling := &node{path: path, value: value}
		n.children = append(n.children, sibling)
	}
}

func insert(path string, value interface{}, current *node) bool {
	pathLength := len(path)
	currentLength := len(current.path)
	common := commonPrefix(path, current.path)
	var found bool

	switch {
	case common == 0:
		found = false

	case common == pathLength:
		// they are the same. Don't insert
		found = true

	case common == currentLength:
		// continue walking down a level with the not matching part
		current.Insert(path[common:], value)
		found = true

	default:
		// set current as the common part and below the new child
		// and the existing children
		child := &node{current.path[common:], current.children, current.value}
		newChild := &node{path: path[common:], value: value}
		current.path = path[:common]
		current.children = []*node{child, newChild}
		current.value = nil
		found = true
	}

	return found
}

func (n *node) Lookup(path string) (interface{}, bool) {
	common := commonPrefix(path, n.path)

	if n.path != "" {
		switch {
		// no match, end this branch search
		case common == 0 || len(n.path) > common:
			return nil, false

		// We found it
		case common == len(path):
			return n.value, true
		}
	}

	for _, child := range n.children {
		r, ok := child.Lookup(path[common:])
		if ok {
			return r, ok
		}
	}

	return nil, false
}

func commonPrefix(a, b string) int {
	i := 0
	m := min(len(a), len(b))
	for ; i < m; i++ {
		if a[i] != b[i] {
			break
		}
	}
	return i
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Debug helpers
// -------------------------
func (n *node) Debug() {
	n.Print(os.Stdout, 0)
}

func (n *node) Format(f fmt.State, c rune) {
	f.Write([]byte(fmt.Sprintf("%s %q", n.path, n.children)))
}

func (n *node) Print(w io.Writer, indent int) {
	if n.path != "" {
		fmt.Fprintf(w, "%s %v\n", format(n.path, indent), n.value)
		indent++
	}

	for _, child := range n.children {
		child.Print(w, indent)
	}
}

func format(str string, length int) string {
	for i := 0; i < length; i++ {
		str = "-" + str
	}
	return str
}
