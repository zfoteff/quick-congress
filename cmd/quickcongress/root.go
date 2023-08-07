package quickcongress

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/zfoteff/quick-congress/pkg/quickcongress/node"
)

// Evaluation root function for the CLI version of the Quick Congress application.
// Continues to prompt for user input until the user quits, an unhandled error is
// encountered, or the program ends naturally with nil node.
func quickCongressCLIEntryPoint(cmd *cobra.Command, args []string) {
	var node node.Node = node.NewHeadMenuNode()

	for {
		if node == nil {
			return
		}

		node = node.Evaluate()
	}
}

// CLI application entry point
func Execute() {
	var rootCmd = &cobra.Command{
		Use:   "quick-congress",
		Short: "quick-congress - a simple CLI to inspect congressional bill/amendments",
		Long:  "Quick Congress: A simple interface for gaining more in-depth knowledge about what the hell is going on in congress",
		Run:   quickCongressCLIEntryPoint,
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error during execution:\n'%s'", err)
		os.Exit(1)
	}
}
