package routers

import (
	"net/http"

	"github.com/guso008/twittorquovadis/bd"
	"github.com/guso008/twittorquovadis/models"
)

/*AltaRelacion realiza el registro de la relacion entre usuarios*/
func AltaRelacion(rw http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(rw, "El parámetro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(t)

	if err != nil {
		http.Error(rw, "Ocurrió un error al intentar insertar relación "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(rw, "Ocurrió un error al intentar insertar relación "+err.Error(), http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}
