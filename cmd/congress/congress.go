package congress

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
	// TODO: abstract with a metadata header that includes previous step for back button, number of selections for the end range

	var menuChoice string

	for {
		print(bin.CongressMenu) // Print menu every loop interation
		fmt.Scanln(&menuChoice)

		// TODO: Check for 'q' character and return -1 to end program

		menuChoiceValue, err := strconv.Atoi(menuChoice)

		if err == nil && menuChoiceValue >= 0 && menuChoiceValue <= 3 {
			return &menuChoiceValue
		} else {
			println("[ERR] Please only enter the options displayed in the menu")
		}
	}
}

func getInputForPastCongressSelection() *int {
	var menuChoice string

	for {
		print(bin.CongressYearSelectionMenu) // Print menu every loop interation
		fmt.Scanln(&menuChoice)

		// TODO: Check for 'b' character and return -2 to return to previous menu

		menuChoiceValue, err := strconv.Atoi(menuChoice)

		if err == nil && menuChoiceValue >= 1 && menuChoiceValue <= 117 {
			return &menuChoiceValue
		} else {
			println("[ERR] Please only enter the options displayed in the menu")
		}
	}
}

// Entry function for the congress CLI menu
func CLIEntryPoint(cmd *cobra.Command, args []string) {
	goEnvErr := godotenv.Load(".env")
	client := client.NewCongressClient(os.Getenv("LIBRARY_OF_CONGRESS_API_KEY"))

	if goEnvErr != nil {
		log.Fatalf("Some error occured. Err: %s", goEnvErr)
	}

	menuChoice := getMenuChoice()

	switch *menuChoice {
	case 0:
		println(cli.GetCurrentCongressSession(client, context.TODO()))
	case 1:
		session := getInputForPastCongressSelection()
		println(cli.GetCongressSession(client, context.TODO(), uint16(*session)))
	case 2:
		println("2")
	case 3:
		println("3")
	default:
		println("[ERR] Please enter one of the menu selections on screen")
	}
}
