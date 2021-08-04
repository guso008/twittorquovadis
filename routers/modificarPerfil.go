package routers

import (
	"encoding/json"
	"net/http"

	"github.com/guso008/twittorquovadis/bd"
	"github.com/guso008/twittorquovadis/models"
)

/*ModificarPerfil modifica el perfil de usuario*/
func ModificarPerfil(rw http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(rw, "Datos Incorrectos"+err.Error(), http.StatusBadRequest)
		return
	}

	var status bool
	status, err = bd.ModificoRegistro(t, IDUsuario)

	if err != nil {
		http.Error(rw, "Ocurrió un error al intentar modificar el registro. Reintente nuevamente"+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(rw, "No se ha logrado modificar el registro del usuario ", http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusCreated)

}
