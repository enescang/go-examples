//Enes Can Güneş - kodlib.com - 2021
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MONGO_URI = "mongodb+srv://user-api:1230321@mongocluster.dnuty.mongodb.net/go_test?retryWrites=true&w=majority"

var (
	client *mongo.Client
	err    error
)

func main() {
	client, err = mongo.NewClient(options.Client().ApplyURI(MONGO_URI))
	if err != nil {
		log.Fatal("Connection Err: ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Create(ctx)
	// ReadAll(ctx)
	// ReadOne()
	// Update("GO")
	// Delete("FORTRAN")

}

func Create(ctx context.Context) {
	postCollection := client.Database("go_test").Collection("posts")

	myPosts := []interface{}{
		bson.D{{Key: "title", Value: "GO"}, {Key: "body", Value: "Hello MongoDB"}},
		bson.D{{Key: "title", Value: "JS"}, {Key: "body", Value: "Hello JavaScript"}},
		bson.D{{Key: "title", Value: "FORTRAN"}, {Key: "body", Value: "Hello FORTRAN"}},
	}

	//To insert one document;
	//use => postCollection.InsertOne(ctx, bson.D{{Key:"the_key", Value:"The value"}})

	result, err := postCollection.InsertMany(ctx, myPosts)
	if err != nil {
		log.Fatal("Insert Many Error: ", err)
	}

	fmt.Println("CREATE OPERATION >>>")
	fmt.Println(result)
	fmt.Println("CREATE OPERATION <<<")
}

func ReadAll(ctx context.Context) {
	postCollection := client.Database("go_test").Collection("posts")

	result, err := postCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer result.Close(ctx)

	var posts []Post
	err = result.All(ctx, &posts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("READ ALL OPERATION >>>")
	fmt.Println(posts)
	fmt.Println("READ ALL OPERATION <<<")
}

func ReadOne() {
	postCollection := client.Database("go_test").Collection("posts")

	var post Post
	filter := bson.D{{Key: "title", Value: "GO"}}
	err = postCollection.FindOne(context.Background(), filter).Decode(&post)
	if err != nil {
		log.Fatal("Error ReadOne ", err)
	}

	fmt.Println("READ ONE OPERATION >>>")
	fmt.Println(post)
	fmt.Println("READ ONE OPERATION <<<")
}

func Update(title string) {
	postCollection := client.Database("go_test").Collection("posts")

	filter := bson.D{{Key: "title", Value: title}}

	newContent := bson.D{{Key: "$set", Value: bson.D{{
		Key:   "body",
		Value: "This body changed!",
	}}}}

	result, err := postCollection.UpdateOne(context.TODO(), filter, newContent)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("UPDATE OPERATION >>>")
	fmt.Println(result.MatchedCount)
	fmt.Println("UPDATE OPERATION <<<")
}

func Delete(title string) {
	postCollection := client.Database("go_test").Collection("posts")

	filter := bson.D{{Key: "title", Value: title}}

	result, err := postCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DELETE OPERATION >>>")
	fmt.Println(result.DeletedCount)
	fmt.Println("DELETE OPERATION <<<")
}

type Post struct {
	Title string `bson:"title,omitempty"`
	Body  string `bson:"body,omitempty"`
}
