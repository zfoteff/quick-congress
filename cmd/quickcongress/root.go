package quickcongress

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/zfoteff/quick-congress/pkg/quickcongress/client"
	"github.com/zfoteff/quick-congress/pkg/quickcongress/controller/cli"
)

const menu string = "Quick Congress"

func congressCLIEntry(cmd *cobra.Command, args []string) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	client := client.NewCongressClient(os.Getenv("LIBRARY_OF_CONGRESS_API_KEY"))
	println(cli.GetCurrentCongressSession(client, context.TODO()))

	switch userIn {
	case 0:

	}
}

func Execute() {
	var rootCmd = &cobra.Command{
		Use:   "quick-congress",
		Short: "quick-congress - a simple CLI to inspect congressional bill/amendments",
		Long:  "Quick Congress: A simple interface for gaining more in-depth knowledge about what the hell is going on in congress",
		Run:   congressCLIEntry,
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error during execution:\n'%s'", err)
		os.Exit(1)
	}
}
