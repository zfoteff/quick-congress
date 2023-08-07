package node

// Interface for CLI menu nodes for the program evaluation loop.
// Nodes will evaluate until a stop node is encountered, or the
// program ends naturally
type Node interface {
	GetNodeInput() int16
	Evaluate() Node
}
