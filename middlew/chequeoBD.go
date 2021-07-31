package middlew

import (
	"net/http"

	"github.com/guso008/twittorquovadis/bd"
)

/*ChequeoBD es el middlew que permite conocer el estado de la BD*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(rw, "Conexión perdida con la Base de Datos", 500)
			return
		}

		next.ServeHTTP(rw, r)
	}
}
