package db

import (
	"database/sql"
	"fmt"

	"github.com/bartektricks/ps-keyboard-admin/internal/model"
	"github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(connStr string) (*Repository, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	return &Repository{db: db}, nil
}

func (r *Repository) Close() error {
	return r.db.Close()
}

func (r *Repository) GetVerificationRequests() ([]model.Request, error) {
	query := `SELECT id, name, "verifiedTags", "notVerifiedTags" FROM games WHERE "notVerifiedTags" IS NOT NULL`

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, fmt.Errorf("error querying database: %w", err)
	}

	defer rows.Close()

	var requests []model.Request

	for rows.Next() {
		var req model.Request
		err := rows.Scan(&req.ID, &req.Name, &req.VerifiedTags, &req.NotVerifiedTags)

		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		requests = append(requests, req)
	}

	return requests, rows.Err()
}

func (r *Repository) AcceptVerification(gameId string) error {
	var tags pq.StringArray

	err := r.db.QueryRow("SELECT \"notVerifiedTags\" FROM games WHERE id = $1", gameId).Scan(&tags)
	if err != nil {
		return fmt.Errorf("error querying database: %w", err)
	}

	_, err = r.db.Exec("UPDATE games SET \"notVerifiedTags\" = NULL, \"verifiedTags\" = $1 WHERE id = $2", tags, gameId)
	if err != nil {
		return fmt.Errorf("error updating database: %w", err)
	}

	return nil
}

func (r *Repository) RejectVerification(gameId string) error {
	_, err := r.db.Exec("UPDATE games SET \"notVerifiedTags\" = NULL WHERE id = $1", gameId)
	if err != nil {
		return fmt.Errorf("error updating database: %w", err)
	}

	return nil
}
