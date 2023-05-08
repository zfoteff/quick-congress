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

var rootCmd = &cobra.Command{
	Use:   "quick-congress",
	Short: "quick-congress - a simple CLI to inspect congressional bill/amendments",
	Long:  "Quick Congress: A simple interface for gaining more in-depth knowledge about what the hell is going on in congress",
	Run: func(cmd *cobra.Command, args []string) {

		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Some error occured. Err: %s", err)
		}

		client := client.NewCongressClient(os.Getenv("LIBRARY_OF_CONGRESS_API_KEY"))
		context := context.TODO()
		println(cli.GetCurrentCongressSession(client, context))
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error during execution:\n'%s'", err)
		os.Exit(1)
	}
}
