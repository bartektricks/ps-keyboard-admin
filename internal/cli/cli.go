package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/bartektricks/ps-keyboard-admin/internal/service"
	"github.com/bartektricks/ps-keyboard-admin/internal/ui"
)

type Flags struct {
	RejectID    string
	AcceptID    string
	Print       bool
	Interactive bool
}

func ParseFlags() *Flags {
	rejectIDFlag := flag.String("reject", "", "ID of the game to reject")
	acceptIDFlag := flag.String("accept", "", "ID of the game to accept")
	printFlag := flag.Bool("print", false, "Print all requests")
	flag.Parse()

	return &Flags{
		RejectID:    *rejectIDFlag,
		AcceptID:    *acceptIDFlag,
		Print:       *printFlag,
		Interactive: len(flag.Args()) == 0,
	}
}

// TODO: use cobra or similar library for better command line parsing
func ExecuteCommand(flags *Flags, verificationService *service.VerificationService) {
	if len(flag.Args()) > 1 {
		ExitOnError(nil, "Only one flag can be used at a time. Use -print, -accept, or -reject.\n")
	}

	if flags.Print {
		err := verificationService.PrintRequests()
		ExitOnError(err, "Error printing requests: %v\n", err)
	}

	if flags.AcceptID != "" {
		err := verificationService.AcceptVerification(flags.AcceptID)
		ExitOnError(err, "Error accepting verification: %v\n", err)

		fmt.Printf("Accepted verification for game ID: %s\n", flags.AcceptID)
	}

	if flags.RejectID != "" {
		err := verificationService.RejectVerification(flags.RejectID)
		ExitOnError(err, "Error rejecting verification: %v\n", err)

		fmt.Printf("Rejected verification for game ID: %s\n", flags.RejectID)
	}

	if flags.Interactive {
		runInteractiveMode(verificationService)
	}
}

func runInteractiveMode(verificationService *service.VerificationService) {
	requests, err := verificationService.GetVerificationRequests()
	ExitOnError(err, "Error getting verification requests: %v\n", err)

	selectedItems, err := ui.RunUI(requests)
	ExitOnError(err, "Error running UI: %v\n", err)

	if len(selectedItems) > 0 {
		fmt.Println("Approved games:")
		for _, item := range selectedItems {
			err := verificationService.AcceptVerification(item.ID)
			if err != nil {
				fmt.Printf("Error accepting verification for game ID %s: %v\n", item.ID, err)
				continue
			}
			fmt.Printf("- %s (ID: %s)\n", item.Name, item.ID)
		}
	} else {
		fmt.Println("No requests were selected.")
	}
}

func ExitOnError(err error, format string, args ...interface{}) {
	if err != nil {
		fmt.Printf(format, args...)
		os.Exit(1)
	}
}
