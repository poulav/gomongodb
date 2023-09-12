package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/poulav/gomongodb/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://poulavbiswas:[PASSWORD_HERE]@cluster0.1ivcbjc.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongo DB Connection Success!")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance created successfully")
}

// Mongo DB Helpers

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie with id", inserted)
}

//Update one value in database

func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified Count", result.ModifiedCount)
}

//Delete one movie in database

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted Movie with Delete Count", deleteCount)
}

// Delete all movies in database
func deleteAllMovies() int64 {
	deleteRes, err := collection.DeleteMany(context.Background(), bson.M{}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted All Records, Delete Count", deleteRes.DeletedCount)
	return deleteRes.DeletedCount
}

// Get All Movies in database
func getAllMovies() []primitive.M {
	allMoviesCur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M
	for allMoviesCur.Next(context.Background()) {
		var movie bson.M
		err := allMoviesCur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer allMoviesCur.Close(context.Background())
	return movies
}

//Controllers ----------------------------------------------------------------

// For Getting All Movies
func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	movies := getAllMovies()

	json.NewEncoder(w).Encode(movies)
}

// Create a new Movie
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	//Accepts the request and decodes it following the structure of Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

// Update a movie
func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	//This gives us the parameters in KeyValue pair format
	params := mux.Vars(r)
	updateOneMovie(params["id"])

	//returns the id of the movie
	json.NewEncoder(w).Encode(params["id"])
}

// Delete a movie
func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

// func Delete All Movies
func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovies()
	json.NewEncoder(w).Encode(count)
}
