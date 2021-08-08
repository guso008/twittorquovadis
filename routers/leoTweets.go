package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/guso008/twittorquovadis/bd"
)

/*LeoTweets leo los tweets*/
func LeoTweets(rw http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(rw, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(rw, "Debe enviar el parametro pagina", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))

	if err != nil {
		http.Error(rw, "Debe enviar el parametro pagina con un valor mayor a 0", http.StatusBadRequest)
	}

	pag := int64(pagina)

	respuesta, correcto := bd.LeoTweets(ID, pag)

	if correcto == false {
		http.Error(rw, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(respuesta)
}
