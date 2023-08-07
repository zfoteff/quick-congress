package node

import (
	"fmt"
	"strconv"
	"strings"
)

// Interface for CLI menu nodes for the program evaluation loop.
// Nodes will evaluate until a stop node is encountered, or the
// program ends naturally
type Node interface {
	GetNodeInput() int8
	Evaluate() *Node
}

func (n *Node) GetNodeInput() int8 {
	var menuChoice string

	for {
		fmt.Print(n.Text)
		fmt.Scanln(&menuChoice)

		//	Check if user inputted a quit command, or a backup command
		switch strings.ToLower(menuChoice) {
		case "q":
			return -1
		case "b":
			return -2
		}

		menuChoiceValue, err := strconv.Atoi(menuChoice)

		if err == nil && menuChoiceValue >= n.StartRange && menuChoiceValue <= n.EndRange {
			return int8(menuChoiceValue)
		} else {
			println("[ERR] Please only enter the options displayed in the menu")
		}
	}
}
