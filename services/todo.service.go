package services

import (
	"context"

	"github.com/rulanugrh/saturnus/entity"
	"github.com/rulanugrh/saturnus/helper"
	"github.com/rulanugrh/saturnus/pb"
	"github.com/rulanugrh/saturnus/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
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
			CreateAt: timestamppb.New(result.CreateAt),
			UpdateAt: timestamppb.New(result.UpdateAt),
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
			CreateAt: timestamppb.New(result.CreateAt),
			UpdateAt: timestamppb.New(result.UpdateAt),
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

	
	result, err := td.repo.Update(req.GetId(), &data)
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
			CreateAt: timestamppb.New(result.CreateAt),
			UpdateAt: timestamppb.New(result.UpdateAt),
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

func (td *TodoServiceServer) FindAll(empty *emptypb.Empty, stream pb.TodoService_FindAllServer) error {
	data, err := td.repo.FindAll()
	if err != nil {
		return helper.PrintError(err)
	}
	
	for _, result := range data {
		stream.Send(&pb.Data{
			Name: result.Name,
			Category: &pb.Category{
				Name: result.Category.Name,
			},
			IsDone: result.IsDone,
			CreateAt: timestamppb.New(result.CreateAt),
			UpdateAt: timestamppb.New(result.UpdateAt),
		})
	}
	
	return nil
}