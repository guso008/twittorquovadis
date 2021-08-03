package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/guso008/twittorquovadis/bd"
	"github.com/guso008/twittorquovadis/models"
)

/*IDUsuario es el ID devuelto del modelo, que se usará en todos los Endpoints */
var Email string

/*ProcesoToken proceso token para extraer sus valores */
var IDUsuario string

/*ProcesoToken proceso token para extraer sus valores*/
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("Quovadis_TwittorGO")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])
	//Validación del TOKEN
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}

		return claims, encontrado, IDUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}

	return claims, false, string(""), err

}
