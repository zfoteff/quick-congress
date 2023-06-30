package model

// Represents a menu option in the quick-congress application
type MenuNode struct {
	Text       string
	StartRange int
	EndRange   int
	Previous   *MenuNode
}

func NewHeadMenuNode(text string, startRange int, endRange int) *MenuNode {
	return &MenuNode{
		Text:       text,
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
