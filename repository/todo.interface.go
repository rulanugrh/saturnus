package repository

import (
	"github.com/rulanugrh/saturnus/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoInterface interface {
	CreateTodo(todo entity.TodoEntity) (entity.TodoEntity, error)
	FindById(id primitive.ObjectID) (entity.TodoEntity, error)
	Update(id primitive.ObjectID, todoUpt entity.TodoEntity) (entity.TodoEntity, error)
	FindAll() ([]entity.TodoEntity, error)
	Delete(id primitive.ObjectID) error
}