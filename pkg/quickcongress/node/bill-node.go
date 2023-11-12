package node

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/zfoteff/quick-congress/bin"
	"github.com/zfoteff/quick-congress/pkg/quickcongress/controller/cli"
)

type BillNode struct {
	MenuNode
	cli.CLIInterface
}

// Create a new bill node that can be evaluated
func NewBillNode(prev *Node) *BillNode {
	return &BillNode{
		MenuNode{
			Text:       bin.BillMenu,
			StartRange: 0,
			EndRange:   2,
			Previous:   prev,
		}, cli.CLIInterface{},
	}
}

func (n *BillNode) GetBillByNumber() *BillNode {
	return nil
}

func (n *BillNode) GetRecentlyUpdatedBills() *BillNode {
	return nil
}

func (n *BillNode) GetNodeInput() int16 {
	var menuChoice string

	for {
		fmt.Print(n.Text)
		fmt.Scanln(&menuChoice)

		// Check if user inputted a quit command, or a backup command
		switch strings.ToLower(menuChoice) {
		case "q":
			// quit
			return -1
		case "b":
			// backup
			return -2
		}

		menuChoiceValue, err := strconv.Atoi(menuChoice)

		if err == nil && menuChoiceValue >= n.StartRange && menuChoiceValue <= n.EndRange {
			return int16(menuChoiceValue)
		} else {
			println("[ERR] Please only enter the options displayed in the menu")
		}
	}
}

// Evaluation function for the congress menu node. Should output to the
// terminal if a command node is reached
func (n *BillNode) Evaluate() Node {
	var nextNode *Node

	switch n.GetNodeInput() {
	case -2:
		nextNode = n.Previous
	case -1:
		nextNode = nil
	case 0:
		nextNode = nil
	}

	return *nextNode
}
