package main

import (
	"log"

	"github.com/guso008/twittorquovadis/bd"
	"github.com/guso008/twittorquovadis/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
