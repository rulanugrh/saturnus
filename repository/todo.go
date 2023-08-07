package repository

import (
	"context"
	"time"

	"github.com/rulanugrh/saturnus/entity"
	"github.com/rulanugrh/saturnus/helper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type todostruct struct {
	todoColl *mongo.Collection
}

func TodoRepository(todoColl *mongo.Collection) TodoInterface {
	return &todostruct{todoColl: todoColl}
}

func (repo *todostruct) CreateTodo(todo *entity.TodoEntity) (*entity.TodoEntity, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	primitive.NewDateTimeFromTime(todo.CreateAt)
	primitive.NewDateTimeFromTime(todo.UpdateAt)
	primitive.NewDateTimeFromTime(todo.DeleteAt)

	res, err := repo.todoColl.InsertOne(ctx, &todo)
	if err != nil {
		return nil, helper.PrintError(err)
	}

	var newTodo entity.TodoEntity
	if errs := repo.todoColl.FindOne(ctx, bson.M{"id": res.InsertedID}).Decode(&newTodo); errs != nil {
		return nil, helper.PrintError(errs)
	}
	return &newTodo, nil
}

func (repo *todostruct) FindById(id string) (*entity.TodoEntity, error) {
	var todo *entity.TodoEntity
	
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	if err := repo.todoColl.FindOne(ctx, bson.M{"_id": objId}).Decode(&todo); err != nil {
		return nil, helper.PrintError(err)
	}

	return todo, nil
}

func (repo *todostruct) FindAll() ([]entity.TodoEntity, error) {
	var todos []entity.TodoEntity
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	result, err := repo.todoColl.Find(ctx, bson.M{})
	if err != nil {
		return []entity.TodoEntity{}, helper.PrintError(err)
	}

	defer result.Close(ctx)
	for result.Next(ctx) {
		var todo entity.TodoEntity
		if errs := result.Decode(&todo); errs != nil {
			return []entity.TodoEntity{}, helper.PrintError(errs)
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

func (repo *todostruct) Update(id string, todoUpt *entity.TodoEntity) (*entity.TodoEntity, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	
	objId, _ := primitive.ObjectIDFromHex(id)

	result := repo.todoColl.FindOneAndUpdate(ctx, bson.M{"_id":objId}, bson.M{"$set": todoUpt})

	var newTodo entity.TodoEntity
	if err := result.Decode(&newTodo); err != nil {
		return nil, helper.PrintError(err)
	}

	return &newTodo, nil
}

func (repo *todostruct) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	
	objId, _ := primitive.ObjectIDFromHex(id)

	_, err := repo.todoColl.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		return err
	}

	return nil
}