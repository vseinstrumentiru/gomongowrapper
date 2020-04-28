package gomongowrapper

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"

	tracewrap "github.com/opencensus-integrations/gomongowrapper"
)

type Database struct {
	db *tracewrap.WrappedDatabase
}

func (wd *Database) Client() *Client {
	cc := wd.db.Client()

	if cc == nil {
		return nil
	}
	return &Client{cc: cc}
}

func (wd *Database) Collection(name string, opts ...*options.CollectionOptions) *Collection {
	if wd.db == nil {
		return nil
	}

	coll := wd.db.Collection(name, opts...)

	if coll == nil {
		return nil
	}

	return &Collection{coll: coll}
}

func (wd *Database) Drop(ctx context.Context) error {
	return handle(wd.db.Drop(ctx))
}

func (wd *Database) ListCollections(ctx context.Context, filter interface{}, opts ...*options.ListCollectionsOptions) (*mongo.Cursor, error) {
	cur, err := wd.db.ListCollections(ctx, filter, opts...)

	return cur, handle(err)
}

func (wd *Database) Name() string                          { return wd.db.Name() }
func (wd *Database) ReadConcern() *readconcern.ReadConcern { return wd.db.ReadConcern() }
func (wd *Database) ReadPreference() *readpref.ReadPref    { return wd.db.ReadPreference() }

func (wd *Database) RunCommand(ctx context.Context, runCommand interface{}, opts ...*options.RunCmdOptions) *mongo.SingleResult {
	res := wd.db.RunCommand(ctx, runCommand, opts...)

	return res
}

func (wd *Database) WriteConcern() *writeconcern.WriteConcern { return wd.db.WriteConcern() }

func (wd *Database) Database() *mongo.Database { return wd.db.Database() }
