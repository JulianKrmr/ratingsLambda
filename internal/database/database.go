package database

import (
	"context"
	"github.com/JulianKrmr/ratingsLambda/domain"
	"github.com/jackc/pgx/v4"
)

func GetConnection() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), "postgres://dajul123:cmgAhE38vpya@ep-proud-fire-860176.eu-central-1.aws.neon.tech/neondb")
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func InsertRatings(db *pgx.Conn, ratings domain.AppRatings) error {
	for _, statement := range ratings.Statements {
		_, err := db.Exec(context.Background(), "INSERT INTO ratings (statements_id, codeword, statement, group_name, rating) VALUES ($1, $2, $3, $4, $5)", ratings.StatementsID, ratings.Codeword, statement.StatementString, statement.Group, statement.Rating)
		if err != nil {
			return err
		}
	}
	return nil
}
