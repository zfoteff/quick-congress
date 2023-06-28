package quickcongress

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/zfoteff/quick-congress/bin"
	"github.com/zfoteff/quick-congress/pkg/quickcongress/client"
	"github.com/zfoteff/quick-congress/pkg/quickcongress/controller/cli"
)

func getMenuChoice() *int {
	var menuChoice string

	for {
		print(bin.MenuString) // Print menu every loop interation
		fmt.Scanln(&menuChoice)
		menuChoiceValue, err := strconv.Atoi(menuChoice)

		if err == nil && menuChoiceValue >= 0 && menuChoiceValue <= 3 {
			return &menuChoiceValue
		} else {
			println("[ERR] Please only enter the options displayed in the menu")
		}
	}
}

func congressCLIEntry(cmd *cobra.Command, args []string) {
	client := client.NewCongressClient(os.Getenv("LIBRARY_OF_CONGRESS_API_KEY"))
	goEnvErr := godotenv.Load(".env")

	if goEnvErr != nil {
		log.Fatalf("Some error occured. Err: %s", goEnvErr)
	}

	menuChoice := getMenuChoice()

	switch *menuChoice {
	case 0:
		println(cli.GetCurrentCongressSession(client, context.TODO()))
	case 1:
		println("1")
	case 2:
		println("2")
	case 3:
		println("3")
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
