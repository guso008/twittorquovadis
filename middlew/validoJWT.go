package middlew

import (
	"net/http"

	"github.com/guso008/twittorquovadis/routers"
)

/*ValidoJWT permite validar el JWT que nos viene en la petición*/
func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization"))

		if err != nil {
			http.Error(rw, "Error en el Token ! "+err.Error(), http.StatusBadRequest)
			return
		}

		next.ServeHTTP(rw, r)
	}
}
