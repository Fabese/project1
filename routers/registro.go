package routers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Fabese/project1/db"
	"github.com/Fabese/project1/models"
)

func Registro(ctx context.Context) models.RespApi {
	var (
		t models.User
		r models.RespApi
	)

	r.Status = 400

	fmt.Println("Entré a Registro")

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		r.Messagge = err.Error()
		fmt.Println(r.Messagge)
		return r
	}
	if len(t.Email) == 0 {
		r.Messagge = "Debe especificar el Email"
		fmt.Println(r.Messagge)
		return r
	}
	if len(t.Password) < 6 {
		r.Messagge = "Debe especificar una contraseña de al menos 6 caracteres"
		fmt.Println(r.Messagge)
		return r
	}

	_, encontrado, _ := db.UserExists(t.Email)
	if encontrado {
		r.Messagge = "Ya existe un usuario registrado con ese email"
		fmt.Println(r.Messagge)
		return r
	}

	_, status, err := db.InsertLog(t)
	if err != nil {
		r.Messagge = "Ocurrio un error al intentar realizar el registro de usuario " + err.Error()
		fmt.Println(r.Messagge)
		return r
	}
	if !status {
		r.Messagge = "No se ha logrado insertar el registro de usuario"
		fmt.Println(r.Messagge)
		return r
	}
	r.Status = 200
	r.Messagge = "Registro OK"

	return r
}
