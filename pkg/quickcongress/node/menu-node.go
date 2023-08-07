package node

import (
	"fmt"
	"github.com/zfoteff/quick-congress/bin"
	"strconv"
)

// Represents a menu in the quick-congress application
type MenuNode struct {
	Text       string
	StartRange int
	EndRange   int
	Previous   *MenuNode
}

func NewHeadMenuNode(startRange int, endRange int) *MenuNode {
	return &MenuNode{
		Text:       bin.AppMenu,
		StartRange: startRange,
		EndRange:   endRange,
		Previous:   nil,
	}
}

func NewMenuNode(text string, startRange int, endRange int, previous *MenuNode) *MenuNode {
	return &MenuNode{
		Text:       text,
		StartRange: startRange,
		EndRange:   endRange,
		Previous:   previous,
	}
}

func StopNode() *MenuNode {
	return &MenuNode{
		Text:       "STOPNODE",
		StartRange: 0,
		EndRange:   0,
		Previous:   nil,
	}
}

func (m *MenuNode) GetNodeInput() int8 {
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
			return int8(menuChoiceValue)
		} else {
			println("[ERR] Please only enter the options displayed in the menu")
		}
	}
}

func (m *MenuNode) Evaluate() *Node {
	switch m.GetNodeInput() {
	case 0:
		// Congress info.
		return NewCongressNode(m)
	case 1:
		// Bill info.
		fmt.Print("Bill info.")
	case 2:
		// Summary info.
		fmt.Print("Bill info.")
	case 3:
		// Representative info.
		fmt.Print("Bill info.")
	}
}
