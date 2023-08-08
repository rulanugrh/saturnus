package configs

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	GRPC struct {
		Host     string
		Port     string
	}
	
	Database struct {
		Host string
		Port string
		User string
		Pass string
		Name string
	}

	HTTP struct {
		Host string
		Port string
	}
}

var app *App

func MongoConn() *mongo.Client {
	conf := GetConfig()
	serverApiOptions := options.ServerAPI(options.ServerAPIVersion1)
	dsn := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority",
		conf.Database.User,
		conf.Database.Pass,
		conf.Database.Host,
	)

	clientOptions := options.Client().ApplyURI(dsn).SetServerAPIOptions(serverApiOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return client

}

var DB *mongo.Client = MongoConn()

func MongoColl(coll string, client *mongo.Client) *mongo.Collection {
	conf := GetConfig()

	return client.Database(conf.Database.Name).Collection(coll)
}

func GetConfig() *App {
	if app == nil {
		app = initConfig()
	}

	return app
}

func initConfig() *App {
	conf := App{}

	if err := godotenv.Load(); err != nil {
		conf.GRPC.Port = "localhost"
		conf.GRPC.Port = ""
		conf.Database.Host = ""
		conf.Database.Port = ""
		conf.Database.User = ""
		conf.Database.Name = ""
		conf.Database.Pass = ""
		conf.HTTP.Port = ""
		conf.HTTP.Host = "localhost"

		return &conf
	}

	conf.GRPC.Host = os.Getenv("GRPC_HOST")
	conf.GRPC.Port = os.Getenv("GRPC_PORT")
	conf.HTTP.Port = os.Getenv("HTTP_PORT")
	conf.HTTP.Host = os.Getenv("HTTP_HOST")
	conf.Database.Host = os.Getenv("MONGODB_HOST")
	conf.Database.Port = os.Getenv("MONGODB_PORT")
	conf.Database.User = os.Getenv("MONGODB_USER")
	conf.Database.Name = os.Getenv("MONGODB_NAME")
	conf.Database.Pass = os.Getenv("MONGODB_PASS")

	return &conf
}