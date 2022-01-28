package mongodb

import (
	"context"
	"reflect"
	"regexp"
	"strings"

	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *Repository) clone(origin interface{}) (clone interface{}, err error) {
	newClone := reflect.New(reflect.TypeOf(origin).Elem()).Interface()
	return newClone, copier.Copy(newClone, origin)
}

func (repo *Repository) makeFilters(filters []string) (bsonFilters bson.M) {
	bsonFilters = bson.M{}
	for _, v := range filters {
		slFilter := strings.Split(v, ":")
		key := slFilter[0]
		operations := slFilter[1]

		switch operations {
		case "ne":
			bsonFilters[key] = bson.M{"$ne": slFilter[2]}
		case "like":
			bsonFilters[key] = bson.M{
				"$regex":   slFilter[2],
				"$options": "i",
			}
		case "eq":
			bsonFilters[key] = slFilter[2]
		case "eqInt":
			bsonFilters[key] = slFilter[2]
		case "isNull":
			bsonFilters[key] = nil
		case "isNotNull":
			bsonFilters[key] = bson.M{"$ne": nil}
		case "id":
			oid, _ := primitive.ObjectIDFromHex(slFilter[2])
			bsonFilters[key] = oid
		default:
			bsonFilters[key] = slFilter[2]
		}
	}

	return bsonFilters
}

func (repo *Repository) makeSorts(sorts []string) (bsonSorts bson.M) {
	bsonSorts = bson.M{}

	for _, v := range sorts {
		slFilter := strings.Split(v, ":")
		field := slFilter[0]
		order := slFilter[1]
		bsonSorts[field] = -1
		if order == "asc" {
			bsonSorts[field] = 1
		}
	}

	return bsonSorts
}

func (repo *Repository) makePagingOpts(page int, perPage int) (opts *options.FindOptions) {
	skip := (page - 1) * perPage
	opts = options.Find()
	opts.SetSkip(int64(skip))

	if perPage > 0 {
		opts.SetLimit(int64(perPage))
	}

	return opts
}

func (repo *Repository) makeSearch(search []string) (bsonSearch bson.A) {
	bsonSearch = bson.A{}
	for _, v := range search {
		bsonItem := bson.M{}
		slSearch := strings.Split(v, ":")
		key := slSearch[0]
		value := slSearch[1]

		bsonItem[key] = bson.M{
			"$regex":   *repo.addBackSlash(&value),
			"$options": "i",
		}

		bsonSearch = append(bsonSearch, bsonItem)
	}
	return bsonSearch
}

func (repo *Repository) addBackSlash(str *string) *string {
	var res string
	reg, err := regexp.Compile("[^A-Za-z0-9]")
	if err != nil {
		return &res
	}
	res = reg.ReplaceAllStringFunc(*str, func(s string) string {
		if s != " " {
			return "\\" + s
		} else {
			return s
		}
	})
	return &res
}

func (repo *Repository) countBson(ctx context.Context, filters bson.M) (total int, err error) {
	cnt, err := repo.Coll.CountDocuments(ctx, filters)
	if err != nil {
		return 0, err
	}
	return int(cnt), nil
}
