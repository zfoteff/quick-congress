package quickcongress

import "github.com/spf13/cobra"

func congressMenu() *cobra.Command {
	var userIn int

	var command *cobra.Command = &cobra.Command{
		Use:   "congress",
		Short: "Get information about a session of congress",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	return command
}
