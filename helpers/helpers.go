package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Mukunth-arya/golangapp/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const db_url = ""
const db_name = "mydb1"
const db_collec = "mycol1"

var collection *mongo.Collection

func init() {
	//client option
	clientOption := options.Client().ApplyURI(db_url)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection = client.Database(db_name).Collection(db_collec)

	//collection instance
	fmt.Println("Collection instance is ready")
}

// MONGODB helpers - file

// insert 1 record
func insertOneData(Data1 models.Data) {
	inserted, err := collection.InsertOne(context.Background(), Data1)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)
}

// update 1 record
func updateOneData(DataId string) {
	id, _ := primitive.ObjectIDFromHex(DataId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"satisfied": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

// delete 1 record
func deleteOneData(DataId string) {
	id, _ := primitive.ObjectIDFromHex(DataId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MOvie got delete with delete count: ", deleteCount)
}

// delete all records from mongodb
func deleteAllData() int64 {

	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("NUmber of movies delete: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// get all movies from database

func getAllDatas() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var Datahold []primitive.M

	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		Datahold = append(Datahold, movie)
	}

	defer cur.Close(context.Background())
	return Datahold
}

// Actual controller - file

func GetMyAllData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	alldatas := getAllDatas()
	json.NewEncoder(w).Encode(alldatas)
}

func CreateData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var Data3 models.Data
	_ = json.NewDecoder(r.Body).Decode(&Data3)
	insertOneData(Data3)
	json.NewEncoder(w).Encode(Data3)

}

func Satisfication(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneData(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneData(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllData()
	json.NewEncoder(w).Encode(count)
}
