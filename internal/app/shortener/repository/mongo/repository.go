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
)

type Repository struct {
	client   *mongo.Client
	uri      string
	name     string
	username string
	password string
	ctx      context.Context
}

func NewRepository(uri, name, username, password string, ctx context.Context) (shortener.RedirectRepository, error) {
	repo := &Repository{
		uri:      uri,
		name:     name,
		username: username,
		password: password,
		ctx:      ctx,
	}

	client, err := newClient(repo)
	if err != nil {
		return nil, errors.Wrap(err, "repository.NewRepository")
	}
	repo.client = client
	return repo, nil
}

func newClient(repo *Repository) (*mongo.Client, error) {

	credential := options.Credential{
		Username: repo.username,
		Password: repo.password,
	}

	// Initialize a new mongo client with options
	client, err := mongo.NewClient(options.Client().ApplyURI(repo.uri).SetAuth(credential))

	err = client.Connect(repo.ctx)

	// Ping MongoDB
	//ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Ping(repo.ctx, readpref.Primary())
	if err != nil {
		return nil, errors.Wrap(err, "repository.newClient")
	} else {
		log.Println("MongoDB Connected!")
	}

	return client, nil
}

func (r Repository) FindByCode(ctx context.Context, code string) (*model.Redirect, error) {
	redirect := &model.Redirect{}
	collection := r.client.Database(r.name).Collection("redirects")
	filter := bson.M{"code": code}
	err := collection.FindOne(ctx, filter).Decode(&redirect)
	if err != nil {
		return nil, errors.Wrap(err, "repository.Redirect.FindByCode")
	}
	return redirect, nil
}

func (r Repository) Save(ctx context.Context, redirect *model.Redirect) (*model.Redirect, error) {
	collection := r.client.Database(r.name).Collection("redirects")
	_, err := collection.InsertOne(ctx, redirect)
	if err != nil {
		errors.Wrap(err, "repository.Redirect.Save")
	}
	return redirect, err
}
