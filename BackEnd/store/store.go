package store

import (
	"context"
	"fmt"
	"reflect"

	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	client   *mongo.Client
	Database *mongo.Database
)

const myDB = "BilibiliCleanPlan"

func init() {
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	Database = client.Database(myDB)
	fmt.Println("Connected to MongoDB")
}

func Insert(colName string, data any) error {
	col := Database.Collection(colName)

	rv := reflect.ValueOf(data)
	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		return fmt.Errorf("data is not a slice or array: %s", rv.Kind())
	}

	var allErr error
	for i := 0; i < rv.Len(); i++ {
		_, err := col.InsertOne(context.TODO(), rv.Index(i).Interface())
		if err != nil {
			allErr = fmt.Errorf("%s\nerror insert data %v to mongodb: %s", allErr, rv.Index(i), err)
		}
	}

	return allErr
}
