package bd

import "golang.org/x/crypto/bcrypt"

/*EncriptarPassword */
func EncriptarPassword(pass string) (string, error) {
	//Cantidad de pasadas para encriptar el password.
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)

	return string(bytes), err
}
