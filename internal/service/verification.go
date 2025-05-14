package service

import (
	"fmt"

	"github.com/bartektricks/ps-keyboard-admin/internal/db"
	"github.com/bartektricks/ps-keyboard-admin/internal/model"
)

type VerificationService struct {
	repo *db.Repository
}

func NewVerificationService(repo *db.Repository) *VerificationService {
	return &VerificationService{repo: repo}
}

func (s *VerificationService) GetVerificationRequests() ([]model.Request, error) {
	return s.repo.GetVerificationRequests()
}

func (s *VerificationService) AcceptVerification(gameId string) error {
	return s.repo.AcceptVerification(gameId)
}

func (s *VerificationService) RejectVerification(gameId string) error {
	return s.repo.RejectVerification(gameId)
}

func (s *VerificationService) PrintRequests() error {
	requests, err := s.GetVerificationRequests()
	if err != nil {
		return err
	}

	fmt.Printf("Found %d requests\n", len(requests))

	for _, req := range requests {
		fmt.Printf("Request ID: %v\n", req.ID)
		fmt.Printf("Request Name: %v\n", req.Name)
		fmt.Printf("Verified Tags: %v\n", req.VerifiedTags)
		fmt.Printf("Not Verified Tags: %v\n", req.NotVerifiedTags)
		fmt.Printf("--------------------\n")
	}

	return nil
}
