// connect databse with mongo DB 
//---------------------------------------------//


package config
import (
"context"
"log"
"os" // import OS for secure //
"time"

"go.mongodb.org/mongo-driver/mongo"
"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database
func ConnectDB() {

// change this line for securing data (mongoDB URI ) //
	uri := os.Getenv("MONGO_URI") 
	if uri == "" {
		log.Fatal("MONGO_URI not found in .env")
	}

	// create context 
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// connect using latest method of mongo db 
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal("DB Error:", err)
	}
	
// ping to check connection 

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("DB Not Responding:", err)
	}

	log.Println("MongoDB Connected")


// select databse 
	DB = client.Database("school")
}