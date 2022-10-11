package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var CUser *mongo.Collection
var CJiu *mongo.Collection
var CBlog *mongo.Collection

func init() {
	LoadTheEnv()
	CreateDBInstance()
}

func LoadTheEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
func CreateDBInstance() {
	connectionString := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	collJiu := os.Getenv("DB_COLLECTION_JIU")
	collUser := os.Getenv("DB_COLLECTION_USER")
	collBlog := os.Getenv("DB_COLLECTION_BLOG")
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	CJiu = client.Database(dbName).Collection(collJiu)
	CUser = client.Database(dbName).Collection(collUser)
	CBlog = client.Database(dbName).Collection(collBlog)
	fmt.Println("Collection instance created!")
}
