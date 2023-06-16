package application

import (
	"context"
	"github.com/JulianKrmr/ratingsLambda/domain"
	"github.com/JulianKrmr/ratingsLambda/internal/database"
	"github.com/sirupsen/logrus"
)

func InsertRatings(ratings domain.AppRatings) (domain.AppRatings, error) {
	db, err := database.GetConnection()
	if err != nil {
		return domain.AppRatings{}, err
	}
	err = db.Ping(context.TODO())
	if err != nil {
		logrus.Error(err)
	}
	defer db.Close(context.Background())

	err = database.InsertRatings(db, ratings)
	if err != nil {
		logrus.Error(err)
		return domain.AppRatings{}, err
	}

	return ratings, nil

}
