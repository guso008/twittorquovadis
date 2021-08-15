package routers

import (
	"net/http"

	"github.com/guso008/twittorquovadis/bd"
)

/*EliminarTweet permite borrar un Tweet determinado*/
func EliminarTweet(rw http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(rw, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	err := bd.BorroTweet(ID, IDUsuario)

	if err != nil {
		http.Error(rw, "Ocurrió un error al intentar borrar el tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)
}
