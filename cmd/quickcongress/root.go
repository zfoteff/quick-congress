package quickcongress

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/zfoteff/quick-congress/bin"
	"github.com/zfoteff/quick-congress/pkg/quickcongress/client"
	"github.com/zfoteff/quick-congress/pkg/quickcongress/controller/cli"
)

func congressCLIEntry(cmd *cobra.Command, args []string) {
	client := client.NewCongressClient(os.Getenv("LIBRARY_OF_CONGRESS_API_KEY"))
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	var menuChoice string
	for {
		print(bin.MenuString)
		fmt.Scan(&menuChoice)
		isNumeric := regexp.MustCompile(`\d`).MatchString(menuChoice)
		if isNumeric {

			if menuChoice < 0 || menuChoice > 3 {
				println("[ERR] Please only enter the options displayed in the menu")
			}
		} else if menuChoice < 0 || menuChoice > 3 {
		} else {
			break
		}
	}

	switch menuChoice {
	case 0:
		println(cli.GetCurrentCongressSession(client, context.TODO()))
	default:
		println("[ERR] Please enter one of the menu selections on screen")
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
