package secretmanager

import (
	"fmt"
	"github.com/Fabese/project1/awsgo"
	"github.com/Fabese/project1/models"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(SecretName string) (models.Secret, error) {
	var datosSecret models.Secret
	fmt.Println("> Pido secreto " + SecretName)
	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	clave, err := svc.GetSecretValue()
	return datosSecret, nil
}
