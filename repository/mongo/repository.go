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
	database   string
	timeout    time.Duration
	credential options.Credential
}

func NewClient(url string, credential options.Credential, timeout int) (*mongo.Client, error) {
	// Initialize a new mongo client with options
	client, err := mongo.NewClient(options.Client().ApplyURI(url).SetAuth(credential))

	// Connect the mongo client to the MongoDB server
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	err = client.Connect(ctx)

	//To close the connection at the end
	defer cancel()

	// Ping MongoDB
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewClient")
	} else {
		log.Println("Connected!")
	}

	return client, nil
}

func NewRepository(url, db string, credential options.Credential, timeout int) (shortener.RedirectRepository, error) {
	repo := &Repository{
		database: db,
	}
	client, err := NewClient(url, credential, timeout)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewRepository")
	}
	repo.client = client
	return repo, nil
}

func (r Repository) FindByCode(code string) (*shortener.Redirect, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	redirect := &shortener.Redirect{}
	collection := r.client.Database(r.database).Collection("redirects")
	filter := bson.M{"code": code}
	err := collection.FindOne(ctx, filter).Decode(&redirect)
	if err != nil {
		return nil, errors.Wrap(err, "repository.Redirect.FindByCode")
	}

	return redirect, nil
}

func (r Repository) FindByID(id uint) (*shortener.Redirect, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	redirect := &shortener.Redirect{}
	collection := r.client.Database(r.database).Collection("redirects")
	filter := bson.M{"id": id}
	err := collection.FindOne(ctx, filter).Decode(&redirect)
	if err != nil {
		return nil, errors.Wrap(err, "repository.Redirect.FindByID")
	}

	return redirect, nil
}

func (r Repository) Save(redirect *shortener.Redirect) (*shortener.Redirect, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection("redirects")
	_, err := collection.InsertOne(ctx,
		bson.M{
			"code":       redirect.Code,
			"url":        redirect.URL,
			"created_at": redirect.CreatedAt,
			"click":      redirect.Click,
		},
	)
	if err != nil {
		errors.Wrap(err, "repository.Redirect.Save")
	}

	return redirect, err
}
