package node

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/zfoteff/quick-congress/bin"
	"github.com/zfoteff/quick-congress/pkg/quickcongress/controller/cli"
)

type CongressNode struct {
	MenuNode
	cli.CLIInterface
}

// Create a new congress node that can be evaluated
func NewCongressNode(prev Node) *CongressNode {
	return &CongressNode{
		MenuNode{
			Text:       bin.CongressMenu,
			StartRange: 0,
			EndRange:   2,
			Previous:   &prev,
		}, cli.CLIInterface{},
	}
}

func (c *CongressNode) GetNodeInput() int16 {
	var menuChoice string

	for {
		fmt.Print(c.Text)
		fmt.Scanln(&menuChoice)

		//	Check if user inputted a quit command, or a backup command
		switch strings.ToLower(menuChoice) {
		case "q":
			// quit
			return -1
		case "b":
			// backup
			return -2
		}

		menuChoiceValue, err := strconv.Atoi(menuChoice)

		if err == nil && menuChoiceValue >= c.StartRange && menuChoiceValue <= c.EndRange {
			return int16(menuChoiceValue)
		} else {
			println("[ERR] Please only enter the options displayed in the menu")
		}
	}
}

// Evaluation function for the congress menu node. Should output to the terminal if a command node is reached
func (c *CongressNode) Evaluate() Node {
	var nextNode Node

	switch c.GetNodeInput() {
	case -2:
		nextNode = *c.Previous
	case -1:
		nextNode = nil
	case 0:
		// Get current congress session and set next node as main menu
		println(c.GetCurrentCongressSession())
		nextNode = NewHeadMenuNode()
	case 1:
		// Get details about a past session of congress
		nextNode = NewSessionNode(c)
	case 2:
		// get range of congress sessions
		println("2")
	}

	return nextNode
}
