package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	ID      primitive.ObjectID `bson:"_id, omitempty" json:"_id, omitempty"`
	Movie   string             `json:"_movie, omitempty"`
	Watched bool               `json:"_watched, omitempty"`
}
