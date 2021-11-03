//https://docs.mongodb.com/drivers/go/current/

package test__123

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Test_mongodb() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second) //必须到达指定的时间，才会调用 cancel
	defer cancel()

	//client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://foo:bar@localhost:27017"))
	//client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://10.2.2.11:27017"))

	if err != nil {

		fmt.Println(err)
		return
	}
	collection := client.Database("test").Collection("bbb")

	type Trainer struct {
		Name string
		Age  int
		City string
	}
	ash := Trainer{"Ash", 10, "Pallet Town"}
	insertResult, err := collection.InsertOne(context.TODO(), ash)
	if err != nil {
		fmt.Println("---------err:")
		fmt.Println(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	////更新
	//filter := bson.D{{"name", "Ash"}}

	//update := bson.D{
	//	{"$inc", bson.D{
	//		{"age", 1},
	//	}},
	//}
	//_, err := collection.UpdateOne(context.TODO(), filter, update)
	//if err != nil {
	//	log.Fatal(err)
	//}

}

/*
err = client.Disconnect(context.TODO())
if err != nil {
    log.Fatal(err)
}
fmt.Println("Connection to MongoDB closed.")

*/
