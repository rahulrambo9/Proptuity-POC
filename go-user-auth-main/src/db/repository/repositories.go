package repository

import (
	"context"
	"fmt"
	"go-user-auth/config"
	entity "go-user-auth/db/entity"
	"go-user-auth/helper"
	"log"
	"net/url"
	"reflect"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	StorageInstance *mongo.Client
	once            sync.Once
)

// Init initializes the MongoDB connection
func Init(cfg config.AccountsConfig) error {
	var err error

	once.Do(func() {
		log.Println("Connecting to MongoDB...")

		// Encode the username and password to be URL-safe
		encodedUsername := url.QueryEscape(cfg.MongoUsername)
		encodedPassword := url.QueryEscape(cfg.MongoPassword)
		var mongoURI string
		if cfg.MongoDBInstanceLocation == "DOCKER" {
			log.Println("Connecting to MongoDB Docker...")
			mongoURI = fmt.Sprintf("mongodb://%s:%s@%s", encodedUsername, encodedPassword, cfg.MongoDBURI)
		} else {
			log.Println("Connecting to MongoDB Atlas...")
			url := fmt.Sprintf("mongodb+srv://%s:%s@%s", encodedUsername, encodedPassword, cfg.MongoDBURISrv)
			log.Println("MongoDB Atlas URL: ", url)
			mongoURI = url
		}
		log.Println("MongoDB URI: ", mongoURI)

		// Set client options
		serverAPI := options.ServerAPI(options.ServerAPIVersion1)
		clientOptions := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

		// Connect to MongoDB
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			panic(err)
		}

		// Check the connection
		// err = client.Ping(context.Background(), nil)
		// if err != nil {
		// 	panic(err)
		// }

		// Send a ping to confirm a successful connection
		if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
			panic(err)
		}

		fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

		log.Println("Connected to MongoDB!")
		StorageInstance = client
	})

	return err
}

// getting database collections
// func GetCollection(cfg config.AccountsConfig, client *mongo.Client, collectionName string) *mongo.Collection {
// 	log.Printf("Using database: %s, collection: %s", cfg.MongoDBName, collectionName)
// 	collection := client.Database(cfg.MongoDBName).Collection(collectionName)
// 	return collection
// }

func InitKeys(cfg config.AccountsConfig) (string, string, error) {
	// Check if the keys are stored in the database
	keysCollection := GetCollection(StorageInstance, cfg.MongoDBName, cfg.MongoEcdsaKeys)
	count, err := keysCollection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		log.Println("Error while checking keys in the database: ", err)
	}

	// If the keys are not stored in the database, generate new keys and store them in the database
	if count == 0 {
		log.Println("Not found the ECDSA keys in DB, Generating new keys...")
		keySecret, keyId, err := helper.GenerateKeyPair()
		if err != nil {
			log.Println("Error while generating keys: ", err)
			return "", "", err
		}

		keys := entity.ServiceKey{
			Id:        primitive.NewObjectID(),
			KeyType:   "ECDSA",
			KeyId:     keyId,
			KeySecret: keySecret,
			CreatedTS: time.Now().Unix(),
		}

		_, err = keysCollection.InsertOne(context.Background(), keys)
		if err != nil {
			log.Println("Error while storing keys in the database: ", err)
			return "", "", err
		}

		log.Println("Keys stored in the database!")
		return keyId, keySecret, nil
	} else {
		// If the keys are stored in the database, retrieve them
		log.Println("Found the ECDSA keys in DB, Fetching keys...")
		var keys entity.ServiceKey
		err := keysCollection.FindOne(context.Background(), bson.M{}).Decode(&keys)
		if err != nil {
			log.Println("Error while fetching keys from the database: ", err)
			return "", "", err
		}

		return keys.KeyId, keys.KeySecret, nil
	}
}

type MongoRepository interface {
	InsertOne(ctx context.Context, collection string, document interface{}) (*mongo.InsertOneResult, error)
	FindOne(ctx context.Context, collection string, filter interface{}, result interface{}) error
	FindMany(ctx context.Context, collection string, filter interface{}, results interface{}, opts ...*options.FindOptions) error
	UpdateOne(ctx context.Context, collection string, filter interface{}, update interface{}) (*mongo.UpdateResult, error)
	DeleteOne(ctx context.Context, collection string, filter interface{}) (*mongo.DeleteResult, error)
}

type mongoRepositoryImpl struct {
	client *mongo.Client
	dbName string
}

func NewMongoRepository(dbName string) MongoRepository {
	return &mongoRepositoryImpl{
		client: StorageInstance,
		dbName: dbName,
	}
}

// InsertOne inserts a document into a specified collection.
func (r *mongoRepositoryImpl) InsertOne(ctx context.Context, collection string, document interface{}) (*mongo.InsertOneResult, error) {
	coll := GetCollection(r.client, r.dbName, collection)
	return coll.InsertOne(ctx, document)
}

// FindOne finds a single document based on the provided filter and returns it as an interface.
// func (r *mongoRepositoryImpl) FindOne(ctx context.Context, collection string, filter interface{}) (interface{}, error) {
// 	coll := GetCollection(r.client, r.dbName, collection)
// 	var result interface{}
// 	err := coll.FindOne(ctx, filter).Decode(&result)
// 	if err != nil {
// 		log.Printf("Error while fetching user: %+v \n", err)
// 		return nil, err
// 	}
// 	return result, nil
// }

func (r *mongoRepositoryImpl) FindOne(ctx context.Context, collection string, filter interface{}, result interface{}) error {
	coll := GetCollection(r.client, r.dbName, collection)
	err := coll.FindOne(ctx, filter).Decode(result)
	if err != nil {
		log.Printf("Error while fetching document: %+v \n", err)
		return err
	}
	return nil
}

// FindMany finds multiple documents based on the provided filter and returns them as an interface.
// func (r *mongoRepositoryImpl) FindMany(ctx context.Context, collection string, filter interface{}, results interface{}, opts ...*options.FindOptions) error {
// 	coll := GetCollection(r.client, r.dbName, collection)
// 	cursor, err := coll.Find(ctx, filter, opts...)
// 	if err != nil {
// 		return err
// 	}
// 	defer cursor.Close(ctx)

// 	for cursor.Next(ctx) {
// 		var result interface{}
// 		if err := cursor.Decode(&result); err != nil {
// 			return err
// 		}
// 		results = append(results, result)
// 	}

// 	return nil
// }

func (r *mongoRepositoryImpl) FindMany(ctx context.Context, collection string, filter interface{}, results interface{}, opts ...*options.FindOptions) error {
	coll := GetCollection(r.client, r.dbName, collection)
	cursor, err := coll.Find(ctx, filter, opts...)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	// Use reflection to dynamically append decoded items to the results slice.
	resultsVal := reflect.ValueOf(results).Elem()
	for cursor.Next(ctx) {
		// Create a new instance of the slice element type to decode into
		elem := reflect.New(resultsVal.Type().Elem()).Interface()
		if err := cursor.Decode(elem); err != nil {
			return err
		}
		// Append the decoded element to the slice
		resultsVal.Set(reflect.Append(resultsVal, reflect.ValueOf(elem).Elem()))
	}

	return nil
}

// UpdateOne updates a single document based on the filter and update.
func (r *mongoRepositoryImpl) UpdateOne(ctx context.Context, collection string, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	coll := GetCollection(r.client, r.dbName, collection)
	return coll.UpdateOne(ctx, filter, update)
}

// DeleteOne deletes a single document based on the provided filter.
func (r *mongoRepositoryImpl) DeleteOne(ctx context.Context, collection string, filter interface{}) (*mongo.DeleteResult, error) {
	coll := GetCollection(r.client, r.dbName, collection)
	return coll.DeleteOne(ctx, filter)
}

// getting database collections
func GetCollection(client *mongo.Client, dbName, collectionName string) *mongo.Collection {
	return client.Database(dbName).Collection(collectionName)
}
