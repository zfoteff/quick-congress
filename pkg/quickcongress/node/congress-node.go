package node

import (
	"github.com/zfoteff/quick-congress/bin"
	"github.com/zfoteff/quick-congress/pkg/quickcongress/controller/cli"
)

type CongressNode struct {
	*MenuNode
}

// Create a new congress node that can be evaluated
func NewCongressNode(prev *MenuNode) *CongressNode {
	return &CongressNode{
		&MenuNode{
			Text:       bin.CongressMenu,
			StartRange: 0,
			EndRange:   2,
			Previous:   prev,
		},
	}
}

func (c *CongressNode) int8() {

}

// Evaluation function for the congress menu node
func (c *CongressNode) Evaluate() *Node {
	switch c.GetNodeInput() {
	case 0:
		// Get current congress session
		cli.GetCurrentCongressSession()
		return NewHeadMenuNode()
	case 1:
		// Get details about a past session of congress
		// session := node.NewMenuNode(bin.CongressYearSelectionMenu, 1, 117)
		// println(cli.GetCongressSession(uint16(*session)))
	case 2:
		// get range of congress sessions
		println("2")
	default:
		println("[ERR] Please enter one of the menu selections on screen")
	}
}
