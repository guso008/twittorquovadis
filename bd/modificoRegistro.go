package bd

import (
	"context"
	"time"

	"github.com/guso008/twittorquovadis/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*ModificoRegistro permite modificar el perfil del usuario*/
func ModificoRegistro(user models.Usuario, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("users")

	registro := make(map[string]interface{})

	if len(user.Nombre) > 0 {
		registro["nombre"] = user.Nombre
	}

	if len(user.Apellidos) > 0 {
		registro["apellidos"] = user.Apellidos
	}
	registro["fechaNacimiento"] = user.FechaNacimiento
	if len(user.Avatar) > 0 {
		registro["avatar"] = user.Avatar
	}

	if len(user.Banner) > 0 {
		registro["banner"] = user.Banner
	}

	if len(user.Biografia) > 0 {
		registro["biografia"] = user.Biografia
	}

	if len(user.Ubicacion) > 0 {
		registro["ubicacion"] = user.Ubicacion
	}

	if len(user.SitioWeb) > 0 {
		registro["sitioWeb"] = user.SitioWeb
	}

	updString := bson.M{
		"$set": registro,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)

	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filtro, updString)
	if err != nil {
		return false, err
	}

	return true, nil
}
