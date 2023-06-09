package quickcongress

import (
	"github.com/spf13/cobra"
	"github.com/zfoteff/quick-congress/bin"
)

func showMenu(cmd *cobra.Command, args []string) {
	println(bin.MenuString)
}

func congressMenu() *cobra.Command {
	// var userIn int

	var command *cobra.Command = &cobra.Command{
		Use:   "congress",
		Short: "Get information about a session of congress",
		Run:   showMenu,
	}

	return command
}
