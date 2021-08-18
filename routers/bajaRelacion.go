package routers

import (
	"net/http"

	"github.com/guso008/twittorquovadis/bd"
	"github.com/guso008/twittorquovadis/models"
)

/*BajaRelacion realiza el borrado de la relacion entre usuarios*/
func BajaRelacion(rw http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(rw, "El parámetro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.BorroRelacion(t)

	if err != nil {
		http.Error(rw, "Ocurrió un error al intentar borrar relación "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(rw, "Ocurrió un error al intentar insertar borrar "+err.Error(), http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}
