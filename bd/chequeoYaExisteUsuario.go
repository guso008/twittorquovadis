package bd

import (
	"context"
	"time"

	"github.com/guso008/twittorquovadis/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ChequeoYaExisteUsuario  recibe un email de párametro y chequea si ya existe en la BD*/
func ChequeoYaExisteUsuario(email string) (models.Uusario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("users")

	condicion := bson.M{"email": email}
	var resultado models.Uusario

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()

	if err != nil {
		return resultado, false, ID
	}

	return resultado, true, ID
}
