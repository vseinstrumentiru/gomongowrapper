package mongodb

import (
	"context"

	tracewrap "github.com/opencensus-integrations/gomongowrapper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type WrappedClient struct {
	cc *tracewrap.WrappedClient
}

func NewClient(opts ...*options.ClientOptions) (*WrappedClient, error) {
	client, err := tracewrap.NewClient(opts...)

	if err != nil {
		return nil, handle(err)
	}

	return &WrappedClient{cc: client}, nil
}

func (wc *WrappedClient) Connect(ctx context.Context) error {
	err := wc.cc.Connect(ctx)

	return handle(err)
}

func (wc *WrappedClient) Database(name string, opts ...*options.DatabaseOptions) *WrappedDatabase {
	db := wc.cc.Database(name, opts...)

	if db == nil {
		return nil
	}

	return &WrappedDatabase{db: db}
}

func (wc *WrappedClient) Disconnect(ctx context.Context) error {
	return handle(wc.cc.Disconnect(ctx))
}

func (wc *WrappedClient) ListDatabaseNames(ctx context.Context, filter interface{}, opts ...*options.ListDatabasesOptions) ([]string, error) {
	dbs, err := wc.cc.ListDatabaseNames(ctx, filter, opts...)

	return dbs, handle(err)
}

func (wc *WrappedClient) ListDatabases(ctx context.Context, filter interface{}, opts ...*options.ListDatabasesOptions) (mongo.ListDatabasesResult, error) {
	dbr, err := wc.cc.ListDatabases(ctx, filter, opts...)

	return dbr, handle(err)
}

func (wc *WrappedClient) Ping(ctx context.Context, rp *readpref.ReadPref) error {
	return handle(wc.cc.Ping(ctx, rp))
}

func (wc *WrappedClient) StartSession(opts ...*options.SessionOptions) (mongo.Session, error) {
	ss, err := wc.cc.StartSession(opts...)
	if err != nil {
		return nil, handle(err)
	}
	return &WrappedSession{Session: ss}, nil
}

func (wc *WrappedClient) UseSession(ctx context.Context, fn func(mongo.SessionContext) error) error {
	return handle(wc.cc.UseSession(ctx, fn))
}

func (wc *WrappedClient) UseSessionWithOptions(ctx context.Context, opts *options.SessionOptions, fn func(mongo.SessionContext) error) error {
	return handle(wc.cc.UseSessionWithOptions(ctx, opts, fn))
}

func (wc *WrappedClient) Client() *mongo.Client { return wc.cc.Client() }
