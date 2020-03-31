package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"

	tracewrap "github.com/opencensus-integrations/gomongowrapper"
)

type WrappedDatabase struct {
	db *tracewrap.WrappedDatabase
}

func (wd *WrappedDatabase) Client() *WrappedClient {
	cc := wd.db.Client()

	if cc == nil {
		return nil
	}
	return &WrappedClient{cc: cc}
}

func (wd *WrappedDatabase) Collection(name string, opts ...*options.CollectionOptions) *WrappedCollection {
	if wd.db == nil {
		return nil
	}

	coll := wd.db.Collection(name, opts...)

	if coll == nil {
		return nil
	}

	return &WrappedCollection{coll: coll}
}

func (wd *WrappedDatabase) Drop(ctx context.Context) error {
	return handle(wd.db.Drop(ctx))
}

func (wd *WrappedDatabase) ListCollections(ctx context.Context, filter interface{}, opts ...*options.ListCollectionsOptions) (*mongo.Cursor, error) {
	cur, err := wd.db.ListCollections(ctx, filter, opts...)

	return cur, handle(err)
}

func (wd *WrappedDatabase) Name() string                          { return wd.db.Name() }
func (wd *WrappedDatabase) ReadConcern() *readconcern.ReadConcern { return wd.db.ReadConcern() }
func (wd *WrappedDatabase) ReadPreference() *readpref.ReadPref    { return wd.db.ReadPreference() }

func (wd *WrappedDatabase) RunCommand(ctx context.Context, runCommand interface{}, opts ...*options.RunCmdOptions) *mongo.SingleResult {
	res := wd.db.RunCommand(ctx, runCommand, opts...)

	return res
}

func (wd *WrappedDatabase) WriteConcern() *writeconcern.WriteConcern { return wd.db.WriteConcern() }

func (wd *WrappedDatabase) Database() *mongo.Database { return wd.db.Database() }
