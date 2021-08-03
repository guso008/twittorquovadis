package routers

import (
	"encoding/json"
	"net/http"

	"github.com/guso008/twittorquovadis/bd"
)

/*VerPerfil permite extraer los valores del perfil*/
func VerPerfil(rw http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(rw, "Debe enviar el parámetro ID", http.StatusBadRequest)
	}

	perfil, err := bd.BuscoPerfil(ID)

	if err != nil {
		http.Error(rw, "Ocurrió un error al intentar buscar el registro "+err.Error(), http.StatusBadRequest)
		return
	}

	rw.Header().Set("context-type", "application/json")
	rw.WriteHeader(http.StatusCreated)

	json.NewEncoder(rw).Encode(perfil)
}
