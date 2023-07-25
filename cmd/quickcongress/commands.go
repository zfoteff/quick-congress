package quickcongress

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zfoteff/quick-congress/bin"
	"github.com/zfoteff/quick-congress/pkg/quickcongress/node"
)

func showMenu(cmd *cobra.Command, args []string) {
	println(bin.CongressMenu)
}

// Get user input for menu node.
// Continuously prompt the
func getMenuNodeInput(node *node.MenuNode) int {
	var menuChoice string

	for {
		fmt.Print(node.Text)
		fmt.Scanln(&menuChoice)

		// Check if user inputted commands to quit the program or go back a node
		switch strings.ToLower(menuChoice) {
		case "q":
			return -1
		case "b":
			return -2
		}

		menuChoiceValue, err := strconv.Atoi(menuChoice)

		if err == nil && menuChoiceValue >= node.StartRange && menuChoiceValue <= node.EndRange {
			return menuChoiceValue
		} else {
			println("[ERR] Please only enter the options displayed in the menu")
		}
	}
}

func quickCongressEntryPoint() *cobra.Command {
	var command *cobra.Command = &cobra.Command{
		Use:   "congress",
		Short: "Get information about a session of congress",
		Run:   showMenu,
	}

	return command
}
