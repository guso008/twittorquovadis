package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/guso008/twittorquovadis/bd"
	"github.com/guso008/twittorquovadis/jwt"
	"github.com/guso008/twittorquovadis/models"
)

/*Login realiza el login*/
func Login(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("content-type", "aplication/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(rw, "Usuario y/o Contraseña inválidos. "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(rw, "El email del usuario es requerido "+err.Error(), 400)
		return
	}

	documento, existe := bd.IntentoLogin(t.Email, t.Password)
	if existe == false {
		http.Error(rw, "Usuario y/o Contraseña inválidos. "+err.Error(), 400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(rw, "Ocurrió un error al intentar generar el Toker correspondiente. "+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)

	json.NewEncoder(rw).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(rw, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
