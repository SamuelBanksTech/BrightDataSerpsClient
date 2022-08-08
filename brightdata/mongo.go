package brightdata

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type BDMongoClient struct {
	mongoOpts  *BrightDataMongoLogOptions
	MongClient *mongo.Client
	Collection *mongo.Collection
	configured bool
}

type mongoLog struct {
	LogLevel  int
	Message   string
	TimeStamp time.Time
	Data      interface{}
}

func NewBDMongoClient(bdmo *BrightDataMongoLogOptions) *BDMongoClient {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)

	clientOptions := options.Client().
		ApplyURI(bdmo.MongoUrl).
		SetServerAPIOptions(serverAPIOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	collection := client.Database(bdmo.Database).Collection(bdmo.Collection)

	return &BDMongoClient{
		mongoOpts:  bdmo,
		MongClient: client,
		Collection: collection,
		configured: true,
	}
}

func (bdmc *BDMongoClient) StoreLog(logLevel int, message string, data interface{}) {

	if bdmc != nil {
		logData := mongoLog{
			LogLevel:  logLevel,
			Message:   message,
			TimeStamp: time.Now(),
			Data:      data,
		}

		if bdmc.configured {
			_, err := bdmc.Collection.InsertOne(context.TODO(), logData)
			if err != nil {
				log.Println("---------- INSERT TO MONGO ERROR -------------")
				log.Println(err)
				return
			}
		}
	}

	return

}
