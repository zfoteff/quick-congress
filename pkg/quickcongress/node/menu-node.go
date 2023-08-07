package node

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/zfoteff/quick-congress/bin"
)

// Represents a menu in the quick-congress application
type MenuNode struct {
	Text       string
	StartRange int
	EndRange   int
	Previous   *Node
}

func NewHeadMenuNode() *MenuNode {
	return &MenuNode{
		Text:       bin.AppMenu,
		StartRange: 0,
		EndRange:   3,
		Previous:   nil,
	}
}

func NewMenuNode(text string, startRange int, endRange int, previous *Node) *MenuNode {
	return &MenuNode{
		Text:       text,
		StartRange: startRange,
		EndRange:   endRange,
		Previous:   previous,
	}
}

func (m *MenuNode) GetNodeInput() int16 {
	var menuChoice string

	for {
		fmt.Print(m.Text)
		fmt.Scanln(&menuChoice)

		//	Check if user inputted a quit command, or a backup command
		switch strings.ToLower(menuChoice) {
		case "q":
			return -1
		case "b":
			return -2
		}

		menuChoiceValue, err := strconv.Atoi(menuChoice)

		if err == nil && menuChoiceValue >= m.StartRange && menuChoiceValue <= m.EndRange {
			return int16(menuChoiceValue)
		} else {
			println("[ERR] Please only enter the options displayed in the menu")
		}
	}
}

func (m *MenuNode) Evaluate() Node {
	var nextNode Node

	switch m.GetNodeInput() {
	case -2:
		// Can't go back on main menu, just exit program
		nextNode = nil
	case -1:
		// Quit program
		nextNode = nil
	case 0:
		// Congress info.
		nextNode = NewCongressNode(m)
	case 1:
		// Bill info.
		fmt.Print("Bill info.")
	case 2:
		// Summary info.
		fmt.Print("Summary info.")
	case 3:
		// Representative info.
		fmt.Print("Representative info.")
	}

	return nextNode
}
