package html

// Node basic html node
type Node interface {

	// Clone clones the current node
	Clone() Node
}
