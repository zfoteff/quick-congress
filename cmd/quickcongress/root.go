package quickcongress

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/zfoteff/quick-congress/bin"
	"github.com/zfoteff/quick-congress/cmd/congress"
	"github.com/zfoteff/quick-congress/pkg/quickcongress/model"
)

// Entry function for the CLI version of the Quick Congress application
// Prompts the user with the main menu, and requests input for a submenu
func quickCongressCLIEntryPoint(cmd *cobra.Command) {
	menuNode := model.NewHeadMenuNode(bin.AppMenu, 0, 3)
	menuSelection := getMenuNodeInput(*&menuNode)

	switch menuSelection {
	case 0:
		congress.CLIEntryPoint(cmd)
	default:
	}
}

// CLI application entry point
func Execute() {
	var rootCmd = &cobra.Command{
		Use:   "quick-congress",
		Short: "quick-congress - a simple CLI to inspect congressional bill/amendments",
		Long:  "Quick Congress: A simple interface for gaining more in-depth knowledge about what the hell is going on in congress",
		Run:   congress.CLIEntryPoint,
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error during execution:\n'%s'", err)
		os.Exit(1)
	}
}
