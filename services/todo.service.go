package services

import (
	"context"

	"github.com/rulanugrh/saturnus/entity"
	"github.com/rulanugrh/saturnus/helper"
	"github.com/rulanugrh/saturnus/pb"
	"github.com/rulanugrh/saturnus/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)
type TodoServiceServer struct {
	pb.UnimplementedTodoServiceServer
	repo repository.TodoInterface
	coll *mongo.Collection
}

func NewGrpcPost(repo repository.TodoInterface, coll *mongo.Collection) *TodoServiceServer {
	todoServ := TodoServiceServer{repo: repo, coll: coll}
	return &todoServ
}

func (td *TodoServiceServer) CreateProduct(ctx context.Context, req *pb.TodoReq) (*pb.TodoRes, error) {
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

	result, err := td.repo.CreateTodo(&data)
	if err != nil {
		response := &pb.TodoRes {
			Code: 500,
			Message: "Failure To Create Todo",
			Data: nil,
		}
		return response, helper.PrintError(err)
	}

	response := &pb.TodoRes {
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

	return response, nil

}

func (td *TodoServiceServer) FindById(ctx context.Context, id *pb.Id) (*pb.TodoRes, error) {
	result, err := td.repo.FindById(id.GetId())
	if err != nil {
		response := &pb.TodoRes {
			Code: 500,
			Message: "Failure Cannot Find Todo By This ID",
			Data: nil,
		}

		return response, helper.PrintError(err)
	}

	
	response := &pb.TodoRes {
		Code: 200,
		Message: "Todo can find",
		Data: &pb.Data{
			Name: result.Name,
			IsDone: result.IsDone,
			Category: &pb.Category{
				Name: result.Category.Name,
			},
		},
	}

	return response, nil
}

func (td *TodoServiceServer) Update(ctx context.Context, req *pb.UpdateTodoReq) (*pb.TodoRes, error) {
	todo := req.Req.GetTodo()
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

	
	result, err := td.repo.Update(req.Id.GetId(), &data)
	if err != nil {
		response := &pb.TodoRes{
			Code: 500,
			Message: "Failure Cannot Find Todo By This ID",
			Data: nil,
		}

		return response, nil
	}

	response := &pb.TodoRes {
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

	return response, nil

}

func (td *TodoServiceServer) Delete(ctx context.Context, id *pb.Id) (*pb.DeleteTodoRes, error) {
	err := td.repo.Delete(id.GetId())
	if err != nil {
		response := &pb.DeleteTodoRes{
			Code: 500,
			Message: "cannot delete todo",
		}

		return response, nil
	}

	response := &pb.DeleteTodoRes{
		Code: 200,
		Message: "success delete todo",
	}

	return response, nil
}

func (td *TodoServiceServer) FindAll(empty *pb.Empty, req pb.TodoService_FindAllServer) error {
	data := entity.TodoEntity{}
	cursor, err := td.coll.Find(context.Background(), bson.M{})
	if err != nil {
		return helper.PrintError(err)
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		err := cursor.Decode(data)
		if err != nil {
			return helper.PrintError(err)
		}

		req.Send(&pb.TodoRes{
			Code: 200,
			Message: "success find all data",
			Data: &pb.Data{
				Name: data.Name,
				IsDone: data.IsDone,
				Category: &pb.Category{
					Name: data.Category.Name,
				},
			},
		})
	}
	if err := cursor.Err(); err != nil {
		return helper.PrintError(err)
	}

	return nil
}