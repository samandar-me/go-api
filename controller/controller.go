package controller

import (
	"awesomeProject/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

const connectionString = "mongodb+srv://Samandar:Samandar@cluster0.7shdabi.mongodb.net/?retryWrites=true&w=majority"
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

func deleteMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount, err := collecton.DeleteOne(context.Background(), filter)
	pErr(err)
	fmt.Println("Movie deleted ", deleteCount)
}

func deleteAllMovie() int64 {
	deleteResult, err := collecton.DeleteMany(context.Background(), bson.D{{}})
	pErr(err)
	fmt.Println("Delete result ", deleteResult)
	return deleteResult.DeletedCount
}

func getAllMovies() []primitive.M {
	cur, err := collecton.Find(context.Background(), bson.D{{}})
	pErr(err)
	var movies []primitive.M
	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		pErr(err)
		movies = append(movies, movie)
	}
	defer cur.Close(context.Background())
	return movies
}

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	var allMovies = getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movie models.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertMovie(movie)
	err := json.NewEncoder(w).Encode(movie)
	if err != nil {
		pErr(err)
	}
}
func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	updateMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	var count = deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}
func pErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
