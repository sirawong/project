package mongodb

import (
	"cinema/config"
	"cinema/entities"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	Client *mongo.Client
	DB     *mongo.Database
	Coll   *mongo.Collection
}

//go:generate mockery --name=CinemaRepository --output=../mocks
type CinemaRepository interface {
	List(ctx context.Context, opt *entities.PageOption, itemType interface{}) (total int, items []interface{}, err error)
	Read(ctx context.Context, filters []string, out interface{}) (err error)
	Create(ctx context.Context, ent interface{}) (ID string, err error)
	Update(ctx context.Context, filters []string, ent interface{}) (err error)
	Delete(ctx context.Context, filters []string) (err error)
	Count(ctx context.Context, filters []string) (total int, err error)
}

func New(dbConn *mongo.Client, appConfig *config.Config) CinemaRepository {
	repo := &Repository{
		Client: dbConn,
		DB:     dbConn.Database(appConfig.MongoDBName),
	}
	repo.Coll = repo.DB.Collection(appConfig.MongoDBCollection)

	return repo
}
