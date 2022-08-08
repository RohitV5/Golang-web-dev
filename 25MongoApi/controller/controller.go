package controller

//learn more about context package
import (
	"context"
	"fmt"
	"log"

	"github.com/rohitv5/Golang-web-dev/25MongoApi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://admin:admin@localhost:27017/?authMechanism=DEFAULT&authSource=admin"

const dbName = "netflix_golang"
const colName = "watchlist"

var collection *mongo.Collection

//connect with mongoDB

//init is special method which run only once initialy
func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection success")

	collection = (*mongo.Collection)(client.Database(dbName).Collection(colName))

	//collection instance
	fmt.Println("Collection instance is ready")

}

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)
}

func updateOneMovie(movieId string, movie model.Netflix) {

	//convert string to mongoDB objectID
	id, _ := primitive.ObjectIDFromHex(movieId)

	// bson.M for shorter and clearer results
	// bson.D similar to above mostly

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Updated 1 movie in db with id: ", result.UpsertedID)
}
