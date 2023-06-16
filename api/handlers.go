package api

import (
	"github.com/JulianKrmr/ratingsLambda/application"
	"github.com/JulianKrmr/ratingsLambda/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Ratings struct {
	Codeword   string      `json:"codeword"`
	Statements []Statement `json:"statements"`
}
type Statement struct {
	StatementString string `json:"statement"`
	Group           string `json:"group"`
	Rating          bool   `json:"rating"`
}

func HandleRatings(ctx *fiber.Ctx) error {
	var ratings Ratings
	if err := ctx.BodyParser(&ratings); err != nil {
		return err
	}
	inserted, err := application.InsertRatings(ratings.transformToDomainRatings())
	if err != nil {
		return err
	}

	return ctx.JSON(inserted)
}

func (r *Ratings) transformToDomainRatings() domain.AppRatings {
	var appRatings domain.AppRatings
	for _, stmt := range r.Statements {
		appRatings.Statements = append(appRatings.Statements, domain.AppStatement{
			StatementString: stmt.StatementString,
			Group:           stmt.Group,
			Rating:          stmt.Rating,
		})
	}
	appRatings.StatementsID = uuid.New()
	appRatings.Codeword = r.Codeword

	return appRatings
}
