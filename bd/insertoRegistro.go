package bd

import (
	"context"
	"time"

	"github.com/guso008/twittorquovadis/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoRegistro es la para final con la BD para insertar los datos del usuario*/
func InsertoRegistro(user models.Uusario) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("users")

	user.Password, _ = EncriptarPassword(user.Password)

	result, err := col.InsertOne(ctx, user)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil

}
