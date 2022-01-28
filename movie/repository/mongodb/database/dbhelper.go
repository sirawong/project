package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//go:generate mockery --name=DatabaseHelper
type DatabaseHelper interface {
	Collection(name string) CollectionHelper
	Client() ClientHelper
}

//go:generate mockery --name=CollectionHelper
type CollectionHelper interface {
	FindOne(context.Context, interface{}) SingleResultHelper
	Find(context.Context, interface{}, MongoOption) (CursorResultHelper, error)
	InsertOne(context.Context, interface{}) (interface{}, error)
	UpdateOne(context.Context, interface{}, interface{}) (interface{}, error)
	DeleteOne(context.Context, interface{}) (interface{}, error)
	CountDocuments(context.Context, interface{}) (int64, error)
}

//go:generate mockery --name=SingleResultHelper
type SingleResultHelper interface {
	Decode(v interface{}) error
}

//go:generate mockery --name=CursorResultHelper
type CursorResultHelper interface {
	Decode(v interface{}) error
	Next(ctx context.Context) bool
	Close(ctx context.Context) error
}

//go:generate mockery --name=ClientHelper
type ClientHelper interface {
	Database(string) DatabaseHelper
	Connect() error
	// StartSession() (mongo.Session, error)
}

type mongoClient struct {
	cl *mongo.Client
}
type mongoDatabase struct {
	db *mongo.Database
}
type mongoCollection struct {
	coll *mongo.Collection
}

type mongoSingleResult struct {
	sr *mongo.SingleResult
}

type mongoCursorResult struct {
	c *mongo.Cursor
}

type mongoSession struct {
	mongo.Session
}

type MongoOption struct {
	*options.FindOptions
}

// func NewClient(cnf *config.Config) (ClientHelper, error) {
// 	c, err := mongo.NewClient(options.Client().SetAuth(
// 		options.Credential{
// 			Username:   cnf.Username,
// 			Password:   cnf.Password,
// 			AuthSource: cnf.DatabaseName,
// 		}).ApplyURI(cnf.Url))

// 	return &mongoClient{cl: c}, err

// }

// func NewDatabase(cnf *config.Config, client ClientHelper) DatabaseHelper {
// 	return client.Database(cnf.DatabaseName)
// }

func (mc *mongoClient) Database(dbName string) DatabaseHelper {
	db := mc.cl.Database(dbName)
	return &mongoDatabase{db: db}
}

func (mc *mongoClient) Connect() error {
	// mongo client does not use context on connect method. There is a ticket
	// with a request to deprecate this functionality and another one with
	// explanation why it could be useful in synchronous requests.
	// https://jira.mongodb.org/browse/GODRIVER-1031
	// https://jira.mongodb.org/browse/GODRIVER-979
	return mc.cl.Connect(nil)
}

func (md *mongoDatabase) Collection(colName string) CollectionHelper {
	collection := md.db.Collection(colName)
	return &mongoCollection{coll: collection}
}

func (md *mongoDatabase) Client() ClientHelper {
	client := md.db.Client()
	return &mongoClient{cl: client}
}

func (mc *mongoCollection) FindOne(ctx context.Context, filter interface{}) SingleResultHelper {
	singleResult := mc.coll.FindOne(ctx, filter)
	return &mongoSingleResult{sr: singleResult}
}

func (mc *mongoCollection) Find(ctx context.Context, filter interface{}, opts MongoOption) (CursorResultHelper, error) {
	cursor, err := mc.coll.Find(ctx, filter, opts.FindOptions)
	return &mongoCursorResult{c: cursor}, err
}

func (mc *mongoCollection) InsertOne(ctx context.Context, document interface{}) (interface{}, error) {
	data, err := mc.coll.InsertOne(ctx, document)
	return data, err
}

func (mc *mongoCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}) (interface{}, error) {
	data, err := mc.coll.UpdateOne(ctx, filter, update)
	return data, err
}

func (mc *mongoCollection) DeleteOne(ctx context.Context, filter interface{}) (interface{}, error) {
	data, err := mc.coll.DeleteOne(ctx, filter)
	return data, err
}

func (mc *mongoCollection) CountDocuments(ctx context.Context, filter interface{}) (int64, error) {
	data, err := mc.coll.CountDocuments(ctx, filter)
	return data, err
}

func (sr *mongoSingleResult) Decode(v interface{}) error {
	return sr.sr.Decode(v)
}

func (c *mongoCursorResult) Decode(v interface{}) error {
	return c.c.Decode(v)
}

func (c *mongoCursorResult) Close(ctx context.Context) error {
	return c.c.Close(ctx)
}

func (c *mongoCursorResult) Next(ctx context.Context) bool {
	return c.c.Next(ctx)
}
