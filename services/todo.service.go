package services

import (
	"context"

	"github.com/rulanugrh/saturnus/entity"
	"github.com/rulanugrh/saturnus/helper"
	"github.com/rulanugrh/saturnus/pb"
	"github.com/rulanugrh/saturnus/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type TodoServiceServer struct {
	pb.UnimplementedTodoServiceServer
	repo repository.TodoInterface
}

func (serv *TodoServiceServer) CreateProduct(ctx context.Context, req *pb.TodoReq) (*pb.TodoRes, error) {
	todo := req.GetTodo()
	data := entity.TodoEntity {
		Name: todo.GetName(),
		IsDone: todo.GetIsDone(),
		Category: entity.CategoryEntity{
			Name: todo.GetCategory().Name,
		},
		CreateAt: todo.GetCreateAt().AsTime(),
		UpdateAt: todo.GetUpdateAt().AsTime(),
		DeleteAt: todo.GetDeleteAt().AsTime(),
	}


	result, err := serv.repo.CreateTodo(data)
	if err != nil {
		response := pb.TodoRes {
			Code: 500,
			Message: "Failure To Create Todo",
			Data: nil,
		}
		return &response, helper.PrintError(err)
	}

	response := pb.TodoRes {
		Code: 200,
		Message: "Success To Create Todo",
		Data: &pb.Data{
			Name: result.Name,
			Category: &pb.Category{
				Name: result.Category.Name,
			},
			IsDone: result.IsDone,
			CreateAt: todo.CreateAt,
			DeleteAt: todo.UpdateAt,
			UpdateAt: todo.UpdateAt,
		},
	}

	return &response, nil

}

func (td *TodoServiceServer) FindById(ctx context.Context, id *pb.Id) (*pb.TodoRes, error) {
	oid, _ := primitive.ObjectIDFromHex(id.GetId())

	todo, err := td.repo.FindById(oid)
	if err != nil {
		response := pb.TodoRes {
			Code: 500,
			Message: "Failure Cannot Find Todo By This ID",
			Data: nil,
		}

		return &response, nil
	}

	response := pb.TodoRes {
		Code: 200,
		Message: "Todo can find",
		Data: &pb.Data{
			Name: todo.Name,
			IsDone: todo.IsDone,
			Category: &pb.Category{
				Name: todo.Category.Name,
			},
		},
	}

	return &response, nil
}