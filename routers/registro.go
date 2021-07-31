package routers

import (
	"encoding/json"
	"net/http"

	"github.com/guso008/twittorquovadis/bd"
	"github.com/guso008/twittorquovadis/models"
)

/*Registro es la funcion para crear en la BD el registro de usuario*/
func Registro(rw http.ResponseWriter, r *http.Request) {
	var t models.Uusario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(rw, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(rw, "El email de usuario es requerido", 400)
		return
	}

	if len(t.Email) < 6 {
		http.Error(rw, "Debe especificar una contraseña mayor de 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(rw, "Ya existe un usuario registrado con ese email: "+t.Email, 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(rw, "Ocurrió un error al intentar realizar el registro de usuario "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(rw, "No se ha logrado insertar el registro del usuario ", 400)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}
