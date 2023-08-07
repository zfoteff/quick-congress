package node

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/zfoteff/quick-congress/bin"
	"github.com/zfoteff/quick-congress/pkg/quickcongress/controller/cli"
)

type SessionNode struct {
	MenuNode
	cli.CLIInterface
}

func NewSessionNode(prev Node) *SessionNode {
	return &SessionNode{
		MenuNode{
			Text:       bin.SessionMenu,
			StartRange: 1,
			EndRange:   118,
			Previous:   &prev,
		}, cli.CLIInterface{},
	}
}

func (s *SessionNode) GetNodeInput() int16 {
	var menuChoice string

	for {
		fmt.Print(s.Text)
		fmt.Scanln(&menuChoice)

		//	Check if user inputted a quit command, or a backup command
		switch strings.ToLower(menuChoice) {
		case "q":
			return -1
		case "b":
			return -2
		}

		menuChoiceValue, err := strconv.Atoi(menuChoice)

		if err == nil && menuChoiceValue >= s.StartRange && menuChoiceValue <= s.EndRange {
			return int16(menuChoiceValue)
		} else {
			println("[ERR] Please only enter the options displayed in the menu")
		}
	}
}

func (s *SessionNode) Evaluate() Node {
	var nextNode Node
	var userChoice = s.GetNodeInput()

	switch userChoice {
	case -2:
		nextNode = *s.Previous
	case -1:
		nextNode = nil
	default:
		// Return response from client and set next node as the main menu node
		nextNode = NewHeadMenuNode()
		println(s.GetCongressSession(uint16(userChoice)))
	}

	return nextNode
}
