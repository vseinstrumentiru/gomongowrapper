package gomongowrapper

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	tracewrap "github.com/opencensus-integrations/gomongowrapper"
)

type Collection struct {
	coll *tracewrap.WrappedCollection
}

func (wc *Collection) Aggregate(ctx context.Context, pipeline interface{}, opts ...*options.AggregateOptions) (*mongo.Cursor, error) {
	cur, err := wc.coll.Aggregate(ctx, pipeline, opts...)

	return cur, handle(err)
}

func (wc *Collection) BulkWrite(ctx context.Context, models []mongo.WriteModel, opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error) {
	bwres, err := wc.coll.BulkWrite(ctx, models, opts...)

	return bwres, handle(err)
}

func (wc *Collection) Clone(opts ...*options.CollectionOptions) (*mongo.Collection, error) {
	col, err := wc.coll.Clone(opts...)

	return col, handle(err)
}

func (wc *Collection) Count(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	count, err := wc.coll.CountDocuments(ctx, filter, opts...)

	return count, handle(err)
}

func (wc *Collection) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	count, err := wc.coll.CountDocuments(ctx, filter, opts...)
	return count, handle(err)
}

func (wc *Collection) Database() *mongo.Database { return wc.coll.Database() }

func (wc *Collection) DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	dmres, err := wc.coll.DeleteMany(ctx, filter, opts...)

	return dmres, handle(err)
}

func (wc *Collection) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	dor, err := wc.coll.DeleteOne(ctx, filter, opts...)

	return dor, handle(err)
}

func (wc *Collection) Distinct(ctx context.Context, fieldName string, filter interface{}, opts ...*options.DistinctOptions) ([]interface{}, error) {
	distinct, err := wc.coll.Distinct(ctx, fieldName, filter, opts...)

	return distinct, handle(err)
}

func (wc *Collection) Drop(ctx context.Context) error {
	return handle(wc.coll.Drop(ctx))
}

func (wc *Collection) EstimatedDocumentCount(ctx context.Context, opts ...*options.EstimatedDocumentCountOptions) (int64, error) {
	count, err := wc.coll.EstimatedDocumentCount(ctx, opts...)

	return count, handle(err)
}

func (wc *Collection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	cur, err := wc.coll.Find(ctx, filter, opts...)

	return cur, handle(err)
}

func (wc *Collection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	return wc.coll.FindOne(ctx, filter, opts...)
}

func (wc *Collection) FindOneAndDelete(ctx context.Context, filter interface{}, opts ...*options.FindOneAndDeleteOptions) *mongo.SingleResult {
	return wc.coll.FindOneAndDelete(ctx, filter, opts...)
}

func (wc *Collection) FindOneAndReplace(ctx context.Context, filter, replacement interface{}, opts ...*options.FindOneAndReplaceOptions) *mongo.SingleResult {
	return wc.coll.FindOneAndReplace(ctx, filter, replacement, opts...)
}

func (wc *Collection) FindOneAndUpdate(ctx context.Context, filter, update interface{}, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult {
	return wc.coll.FindOneAndUpdate(ctx, filter, update, opts...)
}

func (wc *Collection) Indexes() mongo.IndexView { return wc.coll.Indexes() }

func (wc *Collection) InsertMany(ctx context.Context, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	insmres, err := wc.coll.InsertMany(ctx, documents, opts...)

	return insmres, handle(err)
}

func (wc *Collection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	insores, err := wc.coll.InsertOne(ctx, document, opts...)

	return insores, handle(err)
}

func (wc *Collection) Name() string { return wc.coll.Name() }

func (wc *Collection) ReplaceOne(ctx context.Context, filter, replacement interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {
	repres, err := wc.coll.ReplaceOne(ctx, filter, replacement, opts...)

	return repres, handle(err)
}

func (wc *Collection) UpdateMany(ctx context.Context, filter, replacement interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	umres, err := wc.coll.UpdateMany(ctx, filter, replacement, opts...)

	return umres, handle(err)
}

func (wc *Collection) UpdateOne(ctx context.Context, filter, replacement interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	uores, err := wc.coll.UpdateOne(ctx, filter, replacement, opts...)

	return uores, handle(err)
}

func (wc *Collection) Watch(ctx context.Context, pipeline interface{}, opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error) {
	cs, err := wc.coll.Watch(ctx, pipeline, opts...)

	return cs, handle(err)
}

func (wc *Collection) Collection() *mongo.Collection {
	return wc.coll.Collection()
}
