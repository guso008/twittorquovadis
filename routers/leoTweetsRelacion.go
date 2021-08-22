package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/guso008/twittorquovadis/bd"
)

/*LeoTweetsSeguidores lee los tweets de todos nuestros seguidores*/
func LeoTweetsSeguidores(rw http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(rw, "Debe enviar el parametro página", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))

	if err != nil {
		http.Error(rw, "Debe enviar el parametro página como entero mayor a 0", http.StatusBadRequest)
		return
	}

	respuesta, correcto := bd.LeoTweetsSeguidores(IDUsuario, pagina)

	if correcto == false {
		http.Error(rw, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)

	json.NewEncoder(rw).Encode(respuesta)
}
