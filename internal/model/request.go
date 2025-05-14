package model

import "github.com/lib/pq"

type Request struct {
	ID              string           `db:"id"`
	Name            string           `db:"name"`
	VerifiedTags    []pq.StringArray `db:"verifiedTags"`
	NotVerifiedTags []pq.StringArray `db:"notVerifiedTags"`
}
