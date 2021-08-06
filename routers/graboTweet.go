package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/guso008/twittorquovadis/bd"
	"github.com/guso008/twittorquovadis/models"
)

/*GraboTweet grabo tweet en la BD*/
func GraboTweet(rw http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserId:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)

	if err != nil {
		http.Error(rw, "Ocurrió un error al intentar registrar un registro, reitente nuevamente "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(rw, "No se ha logrado insertar el Tweet", http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}
