package main

import (
	"github.com/bartektricks/ps-keyboard-admin/internal/cli"
	"github.com/bartektricks/ps-keyboard-admin/internal/config"
	"github.com/bartektricks/ps-keyboard-admin/internal/db"
	"github.com/bartektricks/ps-keyboard-admin/internal/service"
)

func main() {
	cfg, err := config.Load()
	cli.ExitOnError(err, "Error loading config: %v\n", err)

	repo, err := db.NewRepository(cfg.DatabaseURL)
	cli.ExitOnError(err, "Error creating repository: %v\n", err)
	defer repo.Close()

	verificationService := service.NewVerificationService(repo)

	flags := cli.ParseFlags()
	cli.ExecuteCommand(flags, verificationService)
}
