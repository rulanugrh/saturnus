package helper

import (
	"go.mongodb.org/mongo-driver/mongo"
)


type TodoCollInterface struct {
	coll *mongo.Collection
}
