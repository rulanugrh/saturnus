package repository

import (
	"context"
	"time"

	"github.com/rulanugrh/saturnus/configs"
	"github.com/rulanugrh/saturnus/entity"
	"github.com/rulanugrh/saturnus/helper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type todostruct struct {
	todoColl *mongo.Collection
}

func TodoRepository() TodoInterface {
	var todoColl *mongo.Collection = configs.MongoColl("todo", configs.DB)
	return &todostruct{todoColl: todoColl}
}

func (repo *todostruct) CreateTodo(todo entity.TodoEntity) (entity.TodoEntity, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	primitive.NewDateTimeFromTime(todo.CreateAt)
	primitive.NewDateTimeFromTime(todo.UpdateAt)
	primitive.NewDateTimeFromTime(todo.DeleteAt)

	_, err := repo.todoColl.InsertOne(ctx, &todo)
	if err != nil {
		return entity.TodoEntity{}, helper.PrintError(err)
	}

	return todo, nil
}

func (repo *todostruct) FindById(id string) (entity.TodoEntity, error) {
	var todo entity.TodoEntity
	
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err := repo.todoColl.FindOne(ctx, bson.M{"id":id}).Decode(&todo); err != nil {
		return entity.TodoEntity{}, helper.PrintError(err)
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

func (repo *todostruct) Update(id string, todoUpt entity.TodoEntity) (entity.TodoEntity, error) {
	var todo entity.TodoEntity
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	result, err := repo.todoColl.UpdateOne(ctx, bson.M{"id":id}, bson.M{"$set": todoUpt})
	if err != nil {
		return entity.TodoEntity{}, helper.PrintError(err)
	}

	if result.MatchedCount == 1 {
		if err := repo.todoColl.FindOne(ctx, bson.M{"id": id}).Decode(&todo); err != nil {
			return entity.TodoEntity{}, helper.PrintError(err)
		}
	}

	return todo, nil
}

func (repo *todostruct) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := repo.todoColl.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return err
	}

	return nil
}