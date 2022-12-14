package mongocontroller

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func mongoConnection() *mongo.Client {
	if os.Getenv("GO_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found")
		}
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	return client

}

func GetManyDocuments(db string, collection string) ([]primitive.M, error) {
	client := mongoConnection()

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database(db).Collection(collection)

	filter := bson.D{}
	// opts := options.Find().SetSort(bson.D{{"$natural", -1}}).SetLimit(10)
	opts := options.Find().SetSort(bson.D{{Key: "$natural", Value: -1}}).SetLimit(10)

	cursor, err := coll.Find(context.TODO(), filter, opts)
	if err != nil {
		panic(err)
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results, nil
}

func GetCountDocuments(db string, collection string) (int64, error) {
	client := mongoConnection()

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database(db).Collection(collection)

	filter := bson.D{}

	count, err := coll.CountDocuments(context.TODO(), filter)

	if err != nil {
		panic(err)
	}

	return count, nil
}

func AddOneDocument(document interface{}, db string, collection string) error {
	client := mongoConnection()
	coll := client.Database(db).Collection(collection)

	_, err := coll.InsertOne(context.TODO(), document)

	return err
}
