package db

import "go.mongodb.org/mongo-driver/mongo"

var (
	MongoCN      *mongo.Client
	DatabaseName string
)
