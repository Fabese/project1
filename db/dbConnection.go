package db

import (
	"context"
	"fmt"
	"github.com/Fabese/project1/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoCN      *mongo.Client
	DatabaseName string
)

func DbConnect(ctx context.Context) error {
	user := ctx.Value(models.Key("user")).(string)
	passwd := ctx.Value(models.Key("password")).(string)
	host := ctx.Value(models.Key("host")).(string)
	ConnStr := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", user, passwd, host)

	var clientOptions = options.Client().ApplyURI(ConnStr)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Conexión exitosa con la DB")
	MongoCN = client
	DatabaseName = ctx.Value(models.Key("database")).(string)
	return nil
}

func DbConnected() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	return err == nil
}
