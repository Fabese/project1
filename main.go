package main

import (
	"context"
	"fmt"
	"github.com/Fabese/project1/awsgo"
	"github.com/Fabese/project1/db"
	"github.com/Fabese/project1/handlers"
	"github.com/Fabese/project1/models"
	"github.com/Fabese/project1/secretsmanager"
	e "github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"os"
	"strings"
)

func main() {
	fmt.Println()
	lambda.Start(EjecutoLambda)

}
func EjecutoLambda(ctx context.Context, request e.APIGatewayProxyRequest) (*e.APIGatewayProxyResponse, error) {
	var res *e.APIGatewayProxyResponse
	awsgo.InicializoAWS()
	if !ValidoParametros() {
		res = &e.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en las variables de entorno. Deben incluir 'SecretName', 'BucketName', 'UrlPrefix'",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
	}
	secretModel, err := secretsmanager.GetSecret(os.Getenv("SecretName"))
	if err != nil {
		res = &e.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en la lectura de Secret " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}
	path := strings.Replace(request.PathParameters["Red-social"], os.Getenv("UrlPrefix"), "", -1)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("path"), path)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("user"), secretModel.UserName)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), secretModel.Password)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("host"), secretModel.Host)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("database"), secretModel.Database)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("jwtsign"), secretModel.JWTSign)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("body"), request.Body)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucketName"), os.Getenv("BucketName"))

	err = db.DbConnect(awsgo.Ctx)
	if err != nil {
		res = &e.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error en la conexión a la base de datos " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}
	respApi := handlers.Handlers(awsgo.Ctx, request)
	if respApi.CustomResp == nil {
		res = &e.APIGatewayProxyResponse{
			StatusCode: respApi.Status,
			Body:       respApi.Messagge,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}
	return respApi.CustomResp, nil
}

func ValidoParametros() bool {
	_, traeParametro := os.LookupEnv("SecretName")
	if !traeParametro {
		return traeParametro
	}
	_, traeParametro = os.LookupEnv("BucketName")
	if !traeParametro {
		return traeParametro
	}
	_, traeParametro = os.LookupEnv("UrlPrefix")
	if !traeParametro {
		return traeParametro
	}
	return traeParametro
}
