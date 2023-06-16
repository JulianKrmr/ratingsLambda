package domain

import "github.com/google/uuid"

type AppRatings struct {
	Statements   []AppStatement `json:"statements"`
	StatementsID uuid.UUID
	Codeword     string `json:"codeword"`
}
type AppStatement struct {
	StatementString string `json:"statement"`
	Group           string `json:"group"`
	Rating          bool   `json:"rating"`
}
