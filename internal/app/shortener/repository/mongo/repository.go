package mongo

import (
	"context"
	"github.com/ahmetcancicek/go-url-shortener/internal/app/model"
	"github.com/ahmetcancicek/go-url-shortener/internal/app/shortener"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

type repository struct {
	client   *mongo.Client
	uri      string
	name     string
	username string
	password string
}

func NewRepository(uri, name, username, password string) (shortener.RedirectRepository, *mongo.Client, error) {

	repo := repository{
		uri:      uri,
		name:     name,
		username: username,
		password: password,
	}

	client, err := newClient(uri, username, password)
	if err != nil {
		return nil, client, errors.Wrap(err, "repository.NewRepository")
	}
	repo.client = client

	return repo, client, nil
}

func newClient(uri, username, password string) (*mongo.Client, error) {

	// Credential
	credential := options.Credential{
		Username: username,
		Password: password,
	}

	// Context
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()

	// Initialize a new mongo client with options
	client, err := mongo.NewClient(options.Client().ApplyURI(uri).SetAuth(credential))

	// Connect
	err = client.Connect(ctx)

	// Ping MongoDB
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, errors.Wrap(err, "repository.newClient")
	} else {
		log.Println("MongoDB Connected!")
	}

	return client, nil
}

func (r repository) FindByCode(ctx context.Context, code string) (*model.Redirect, error) {
	redirect := &model.Redirect{}
	collection := r.client.Database(r.name).Collection("redirects")
	filter := bson.M{"code": code}
	err := collection.FindOne(ctx, filter).Decode(&redirect)
	if err != nil {
		return nil, errors.Wrap(err, "repository.Redirect.FindByCode")
	}
	return redirect, nil
}

func (r repository) Save(ctx context.Context, redirect *model.Redirect) (*model.Redirect, error) {
	collection := r.client.Database(r.name).Collection("redirects")
	_, err := collection.InsertOne(ctx, redirect)
	if err != nil {
		errors.Wrap(err, "repository.Redirect.Save")
	}
	return redirect, err
}
