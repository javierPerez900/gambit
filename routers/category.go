package routers

import (
	"encoding/json"
	"strconv"

	// "github.com/aws/aws-lambda-go/events"
	"github.com/javier/gambit/bd"
	"github.com/javier/gambit/models"
)

func InsertCategory(body string, User string) (int, string) {
	var t models.Category

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	if len(t.CategName) == 0 {
		return 400, "Debe especificar el Nombre (Title) de la categoría"
	}

	if len(t.CategPath) == 0 {
		return 400, "Debe especificar el Path (Ruta) de la categoría"
	}

	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	result, err2 := bd.InsertCategory(t)
	if err2 != nil {
		return 400, "Ocurrió un error al intentar realizar el registro de la categoría " + t.CategName + " > " + err2.Error()
	}

	return 200, "{ CategID: " + strconv.Itoa(int(result)) + "}"
}

func UpdateCategory(body string, User string, id int) (int, string){
	var t models.Category

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos" + err.Error()
	} 

	if len(t.CategName) == 0 && len(t.CategPath) == 0 {
		return 400, "Debe especificar CategName y CategPath para actualizar"
	}

	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	t.CategID=id
	err2 := bd.UpdateCategory(t)
	if err2 != nil {
		return 400, "Ocurrió un error al intentar realizar el UPDATE de la categoría " + strconv.Itoa(id) + " > " + err.Error()
	}

	return 200, "Update OK"
}