package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/guso008/twittorquovadis/bd"
)

/*ObtenerBanner envía el Avatar HTTP*/
func ObtenerBanner(rw http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(rw, "Debe enviar el párametro ID! ", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)

	if err != nil {
		http.Error(rw, "Usuario no encontrado! ", http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uploads/banners" + perfil.Banner)

	if err != nil {
		http.Error(rw, "Imagen no encontrada! ", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(rw, OpenFile)
	if err != nil {
		http.Error(rw, "Error al copiar la imagen! ", http.StatusBadRequest)
		return
	}
}
