package node

// Interface for CLI menu nodes for the program evaluation loop.
// Nodes will evaluate until a stop node is encountered, or the
// program ends naturally
type nodeRead interface {
	GetNodeInput() int8
}

type nodeEvaluate interface {
	Evaluate() *Node
}

type Node struct {
	nodeRead
	nodeEvaluate
	text string

	prev *Menu
}

func (n *Node) GetNodeInput() int8 {
	var menuChoice string

	for {
		fmt.Print(node.Text)
		fmt.Scanln(&menuChoice)

		//	Check if user inputted a quit command, or a backup command
		switch strings.ToLower(menuChoice) {
		case "q":
			return -1
		case "b":
			return -2
		}

		menuChoiceValue, err := strconv.Atoi(menuChoice)

		if err == nil && menuChoiceValue >= node.StartRange && menuChoiceValue <= node.EndRange {
			return int8(menuChoiceValue)
		} else {
			println("[ERR] Please only enter the options displayed in the menu")
		}
	}
}
