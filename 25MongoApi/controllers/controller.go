package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	model "github.com/Aryan0404/mongoapi/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://aryanpoojary04:aryan1234@cluster0.hnlcvmt.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName = "netflix"
const colName = "watchList"

// Most important
var collection *mongo.Collection
//connect with mongoDB
func init(){
	//clientOption
	clientOption:=options.Client().ApplyURI(connectionString)
	//connect to mongodb
	client,err:=mongo.Connect(context.TODO(),clientOption)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Mongodb connection success")
	collection=client.Database(dbName).Collection(colName)
	//collection instance
	fmt.Println("Collection instance is ready")
}
func insertOneMovie(movie model.Netflix){
	inserted,err:=collection.InsertOne(context.Background(),movie)
	if err!=nil{
		log.Fatal(err)

	}
	fmt.Println("Inserted 1 movie in db with id:",inserted.InsertedID)
}
func updateOneMovie(movieId string){
	id,_:=primitive.ObjectIDFromHex(movieId)
	filter:=bson.M{"_id":id}
	update:=bson.M{"$set":bson.M{"watched":true}}
	result,err:=collection.UpdateOne(context.Background(),filter,update)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Modified count is :",result.ModifiedCount)
}

func deleteOneMovie(movieId string){
	id,_:=primitive.ObjectIDFromHex(movieId)
	filter:=bson.M{"_id":id}
	deleted,err:=collection.DeleteOne(context.Background(),filter)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("The deletedCount is",deleted.DeletedCount)
}
func deleteAllMovie() int64{

	deleted,err:=collection.DeleteMany(context.Background(),bson.D{{}},nil)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("The deleted amount is ",deleted.DeletedCount)
	return deleted.DeletedCount
}
func getAllMovies() []primitive.M{
	cur,err:=collection.Find(context.Background(),bson.D{{}})
	if err!=nil{
		log.Fatal(err)

	}
	var movies []primitive.M
	for cur.Next(context.Background()){
		var movie bson.M
		err:=cur.Decode(&movie)
		if err!=nil{
			log.Fatal(err)
		}
		movies=append(movies,movie)

	}
	defer cur.Close(context.Background())

	return movies
}
func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","POST")

	var movie model.Netflix
	_=json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)

}
func MarkAsWatched(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","PUT")
	params:=mux.Vars(r)	
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
func DeleteAMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")
	params:=mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}
func DeleteAllMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")
	count:=deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}