package controller

//learn more about context package
import (
	"context"
	"fmt"
	"log"

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
	client,err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection success")

	collection = (*mongo.Collection)(client.Database(dbName).Collection(colName))

	//collection instance
	fmt.Println("Collection instance is ready")

	

}
