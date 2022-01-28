package mongodb

import (
	"context"
	"user/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *repository) List(ctx context.Context, opt *entities.PageOption, itemType interface{}) (total int, items []interface{}, err error) {
	var filters bson.M
	var opts *options.FindOptions
	if opt != nil {
		opts = repo.makePagingOpts(opt.Page, opt.PerPage)
		if opt.Filters != nil && len(opt.Filters) > 0 && opt.Search != nil && len(opt.Search) > 0 {
			filters = bson.M{
				"$and": bson.A{
					repo.makeFilters(opt.Filters),
				},
				"$or": repo.makeSearch(opt.Search),
			}
		} else if opt.Filters != nil && len(opt.Filters) > 0 {
			filters = repo.makeFilters(opt.Filters)
		} else if opt.Search != nil && len(opt.Search) > 0 {
			filters = bson.M{
				"$or": repo.makeSearch(opt.Search),
			}
		}

		if opt.Sorts != nil && len(opt.Sorts) > 0 {
			opts.Sort = repo.makeSorts(opt.Sorts)
		}
	}

	total, err = repo.countBson(ctx, filters)
	if err != nil {
		return 0, nil, err
	}

	cursor, err := repo.Coll.Find(ctx, filters, opts)
	if err != nil {
		return 0, nil, err
	}
	defer func() { _ = cursor.Close(ctx) }()

	for cursor.Next(ctx) {
		item, err := repo.clone(itemType)
		if err != nil {
			return 0, nil, err
		}
		err = cursor.Decode(item)
		if err != nil {
			return 0, nil, err
		}
		items = append(items, item)
	}

	return total, items, nil
}

func (repo *repository) Read(ctx context.Context, filters []string, out interface{}) (err error) {
	conditions := repo.makeFilters(filters)
	return repo.Coll.FindOne(ctx, conditions).Decode(out)
}

func (repo *repository) Create(ctx context.Context, ent interface{}) (ID string, err error) {
	inserted, err := repo.Coll.InsertOne(ctx, ent)
	if err != nil {
		return "", err
	}
	return inserted.InsertedID.(string), nil
}

func (repo *repository) Update(ctx context.Context, filters []string, ent interface{}) (err error) {
	conditions := repo.makeFilters(filters)
	_, err = repo.Coll.UpdateOne(ctx, conditions, bson.M{"$set": ent})
	if err != nil {
		return err
	}
	return nil
}

func (repo *repository) Delete(ctx context.Context, filters []string) (err error) {
	conditions := repo.makeFilters(filters)
	_, err = repo.Coll.DeleteOne(ctx, conditions)
	if err != nil {
		return err
	}
	return nil
}

func (repo *repository) Count(ctx context.Context, filters []string) (total int, err error) {
	cnt, err := repo.Coll.CountDocuments(ctx, repo.makeFilters(filters))
	if err != nil {
		return 0, err
	}
	return int(cnt), nil
}
