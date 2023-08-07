package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoEntity struct {
	ID primitive.ObjectID `json:"id,omitempty" bson:"id,omitempty"`
	Name string `json:"name" bson:"name"`
	IsDone bool `json:"is_done" bson:"is_done"`
	Category CategoryEntity `json:"category" bson:"category"`
	CreateAt time.Time `json:"create_at" bson:"create_at"`
	DeleteAt time.Time `json:"delete_at" bson:"delete_at"`
	UpdateAt time.Time `json:"update_at" bson:"update_at"`
}

type CategoryEntity struct {
	ID primitive.ObjectID `json:"id,omitempty" bson:"id,omitempty"`
	Name string `json:"name" bson:"name"`
}