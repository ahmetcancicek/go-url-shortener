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

type Repository struct {
	client   *mongo.Client
	uri      string
	name     string
	username string
	password string
	timeout  time.Duration
}

func NewRepository(uri, name, username, password string, tiomeout int) (shortener.RedirectRepository, error) {
	repo := &Repository{
		uri:      uri,
		name:     name,
		username: username,
		password: password,
		timeout:  time.Duration(tiomeout),
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

func (r Repository) FindByCode(code string) (*model.Redirect, error) {
	redirect := &model.Redirect{}
	collection := r.client.Database(r.name).Collection("redirects")
	filter := bson.M{"code": code}
	err := collection.FindOne(context.TODO(), filter).Decode(&redirect)
	if err != nil {
		return nil, errors.Wrap(err, "repository.Redirect.FindByCode")
	}
	return redirect, nil
}

func (r Repository) Save(redirect *model.Redirect) (*model.Redirect, error) {
	collection := r.client.Database(r.name).Collection("redirects")
	_, err := collection.InsertOne(context.TODO(), redirect)
	if err != nil {
		errors.Wrap(err, "repository.Redirect.Save")
	}
	return redirect, err
}
