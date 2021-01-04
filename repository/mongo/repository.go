package mongo

import (
	"context"
	"github.com/ahmetcancicek/go-url-shortener/shortener"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"log"
	"time"
)

type Repository struct {
	client   *mongo.Client
	database string
}

func NewClient(url string) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client, nil
}

func NewRepository(url, db string) (shortener.RedirectRepository, error) {
	repo := &Repository{
		database: db,
	}
	client, err := NewClient(url)
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

func (r Repository) FindByID(id string) (*shortener.Redirect, error) {
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Generate ID
	redirect.ID, _ = uuid.New()

	// Save
	collection := r.client.Database(r.database).Collection("redirects")
	_, err := collection.InsertOne(ctx, bson.M{
		"ID":         redirect.ID,
		"code":       redirect.Code,
		"url":        redirect.URL,
		"click":      redirect.Click,
		"created_at": redirect.CreatedAt,
	})
	if err != nil {
		errors.Wrap(err, "repository.Redirect.Save")
	}

	return redirect, err
}

func (r Repository) Update(redirect *shortener.Redirect) (*shortener.Redirect, error) {
	// TODO: Implement this function
	panic("implement me")
}
