package db

import (
	"context"
	"github.com/Fabese/project1/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertLog(u models.User) (string, bool, error) {
	ctx := context.TODO()

	database := MongoCN.Database(DatabaseName)
	col := database.Collection("users")

	u.Password, _ = EncryptPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil

}
