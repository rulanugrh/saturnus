package repository

import (
	"github.com/rulanugrh/saturnus/entity"
)

type TodoInterface interface {
	CreateTodo(todo *entity.TodoEntity) (*entity.TodoEntity, error)
	FindById(id string) (*entity.TodoEntity, error)
	Update(id string, todoUpt *entity.TodoEntity) (*entity.TodoEntity, error)
	Delete(id string) error
}