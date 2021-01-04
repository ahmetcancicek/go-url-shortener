package mongo

import (
	"context"
	"github.com/ahmetcancicek/go-url-shortener/shortener"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
