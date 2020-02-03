package tree

// Node represents the root of a sub tree
type Node struct {
	name     string
	children []*Node
}

// MakeNode makes a new node for a tree
func MakeNode(name string) *Node {
	tree := Node{name: name, children: []*Node{}}
	return &tree
}

// IsLeaf returns true if this Node has no children
func (node *Node) IsLeaf() bool {
	return len(node.children) == 0
}

// Append adds the provided child as a dependent of this Node
func (node *Node) Append(child *Node) {
	node.children = append(node.children, child)
}

// GetPaths returns a slice of strings of the paths from each node to root
func (node *Node) GetPaths(prefix string) []string {
	var pathToRoot string
	results := []string{}

	if len(prefix) > 0 {
		pathToRoot = prefix + "->" + node.name
		results = append(results, pathToRoot)
	} else {
		pathToRoot = node.name
	}

	for _, child := range node.children {
		results = append(results, child.GetPaths(pathToRoot)...)
	}
	return results
}
