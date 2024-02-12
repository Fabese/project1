package db

import (
	"context"
	"github.com/Fabese/project1/models"
	"go.mongodb.org/mongo-driver/bson"
)

func UserExists(email string) (models.User, bool, string) {
	ctx := context.TODO()

	database := MongoCN.Database(DatabaseName)
	col := database.Collection("users")

	condition := bson.M{"email": email}

	var resultado models.User

	err := col.FindOne(ctx, condition)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
