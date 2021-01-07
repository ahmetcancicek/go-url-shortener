package mongo

import (
	"context"
	"github.com/ahmetcancicek/go-url-shortener/shortener"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

type Repository struct {
	client     *mongo.Client
	dbURL      string
	dbName     string
	timeout    time.Duration
	credential options.Credential
}

func NewRepository(dbURL, username, password, dbName string, timeout int) (shortener.RedirectRepository, error) {
	repo := &Repository{
		dbURL:   dbURL,
		dbName:  dbName,
		timeout: time.Duration(timeout),
		credential: options.Credential{
			Username: username,
			Password: password,
		},
	}

	client, err := newClient(repo)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewRepository")
	}
	repo.client = client
	return repo, nil
}

func newClient(repo *Repository) (*mongo.Client, error) {
	// Initialize a new mongo client with options
	client, err := mongo.NewClient(options.Client().ApplyURI(repo.dbURL).SetAuth(repo.credential))

	// Connect the mongo client to the MongoDB server
	ctx, cancel := context.WithTimeout(context.Background(), repo.timeout)
	err = client.Connect(ctx)

	//To close the connection at the end
	defer cancel()

	// Ping MongoDB
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return nil, errors.Wrap(err, "repository.newClient")
	} else {
		log.Println("MongoDB Connected!")
	}

	return client, nil
}

func (r Repository) FindByCode(code string) (*shortener.Redirect, error) {
	redirect := &shortener.Redirect{}
	collection := r.client.Database(r.dbName).Collection("redirects")
	filter := bson.M{"code": code}
	err := collection.FindOne(context.TODO(), filter).Decode(&redirect)
	if err != nil {
		return nil, errors.Wrap(err, "repository.Redirect.FindByCode")
	}
	return redirect, nil
}

func (r Repository) Save(redirect *shortener.Redirect) (*shortener.Redirect, error) {
	collection := r.client.Database(r.dbName).Collection("redirects")
	_, err := collection.InsertOne(context.TODO(), redirect)
	if err != nil {
		errors.Wrap(err, "repository.Redirect.Save")
	}
	return redirect, err
}
