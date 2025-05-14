package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bartektricks/ps-keyboard-admin/internal/config"
	"github.com/bartektricks/ps-keyboard-admin/internal/db"
	"github.com/bartektricks/ps-keyboard-admin/internal/service"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}

	rejectIDFlag := flag.String("reject", "", "ID of the game to reject")
	acceptIDFlag := flag.String("accept", "", "ID of the game to accept")
	printFlag := flag.Bool("print", false, "Print all requests")
	flag.Parse()

	repo, err := db.NewRepository(cfg.DatabaseURL)
	if err != nil {
		fmt.Printf("Error connecting to the database: %v\n", err)
		os.Exit(1)
	}
	defer repo.Close()

	verificationService := service.NewVerificationService(repo)

	if *printFlag {
		if err := verificationService.PrintRequests(); err != nil {
			fmt.Printf("Error printing requests: %v\n", err)
			os.Exit(1)
		}
	}

	if *acceptIDFlag != "" {
		if err := verificationService.AcceptVerification(*acceptIDFlag); err != nil {
			fmt.Printf("Error accepting verification: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Accepted verification for game ID: %s\n", *acceptIDFlag)
	}

	if *rejectIDFlag != "" {
		if err := verificationService.RejectVerification(*rejectIDFlag); err != nil {
			fmt.Printf("Error rejecting verification: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Rejected verification for game ID: %s\n", *rejectIDFlag)
	}
}
