package repository

import (
	"github.com/rulanugrh/saturnus/entity"
)

type TodoInterface interface {
	CreateTodo(todo *entity.TodoEntityReq) (*entity.TodoEntityDB, error)
	FindById(id string) (*entity.TodoEntityDB, error)
	Update(id string, todoUpt entity.TodoEntityReq) (entity.TodoEntityDB, error)
	FindAll() ([]entity.TodoEntityDB, error)
	Delete(id string) error
}