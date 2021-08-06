package models

import "time"

/*GraboTweet es el formato o esctructura que tendrá nuestro Tweet en la BD*/
type GraboTweet struct {
	UserId  string    `bson:"userid" json:"userid, omitempty"`
	Mensaje string    `bson:"mensaje" json:"mensaje, omitempty"`
	Fecha   time.Time `bson:"fecha" json:"fecha, omitempty"`
}
