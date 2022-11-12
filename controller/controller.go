package controller

import (
	"awesomeProject/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const connectionString = "mongodb+srv://Samandar:Coder#2003@cluster0.7shdabi.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchList"

var collecton *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)
	pErr(err)

	collecton = client.Database(dbName).Collection(colName)
}

func insertMovie(movie models.Netflix) {
	inserted, err := collecton.InsertOne(context.Background(), movie)
	pErr(err)
	fmt.Println("Inserted movie", inserted.InsertedID)
}
func updateMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := collecton.UpdateOne(context.Background(), filter, update)
	pErr(err)
	fmt.Println(result.ModifiedCount)
}
func pErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
