package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name         string             `bson:"name" json:"name,omitempty"`
	LastName     string             `bson:"lastName" json:"lastName,omitempty"`
	BirthDate    time.Time          `bson:"birthDate" json:"birthDate,omitempty"`
	Email        string             `bson:"email" json:"email,omitempty"`
	Password     string             `bson:"password" json:"password,omitempty"`
	Avatar       string             `bson:"avatar" json:"avatar,omitempty"`
	Banner       string             `bson:"banner" json:"banner,omitempty"`
	Biography    string             `bson:"biography" json:"biography,omitempty"`
	Localization string             `bson:"localization" json:"localization,omitempty"`
	Website      string             `bson:"website" json:"website,omitempty"`
}


