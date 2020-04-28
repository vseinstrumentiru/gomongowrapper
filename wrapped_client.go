package gomongowrapper

import (
	"context"

	tracewrap "github.com/opencensus-integrations/gomongowrapper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Client struct {
	cc *tracewrap.WrappedClient
}

func NewClient(opts ...*options.ClientOptions) (*Client, error) {
	client, err := tracewrap.NewClient(opts...)

	if err != nil {
		return nil, handle(err)
	}

	return &Client{cc: client}, nil
}

func (wc *Client) Connect(ctx context.Context) error {
	err := wc.cc.Connect(ctx)

	return handle(err)
}

func (wc *Client) Database(name string, opts ...*options.DatabaseOptions) *Database {
	db := wc.cc.Database(name, opts...)

	if db == nil {
		return nil
	}

	return &Database{db: db}
}

func (wc *Client) Disconnect(ctx context.Context) error {
	return handle(wc.cc.Disconnect(ctx))
}

func (wc *Client) ListDatabaseNames(ctx context.Context, filter interface{}, opts ...*options.ListDatabasesOptions) ([]string, error) {
	dbs, err := wc.cc.ListDatabaseNames(ctx, filter, opts...)

	return dbs, handle(err)
}

func (wc *Client) ListDatabases(ctx context.Context, filter interface{}, opts ...*options.ListDatabasesOptions) (mongo.ListDatabasesResult, error) {
	dbr, err := wc.cc.ListDatabases(ctx, filter, opts...)

	return dbr, handle(err)
}

func (wc *Client) Ping(ctx context.Context, rp *readpref.ReadPref) error {
	return handle(wc.cc.Ping(ctx, rp))
}

func (wc *Client) PingContext(ctx context.Context) error {
	return wc.Ping(ctx, nil)
}

func (wc *Client) StartSession(opts ...*options.SessionOptions) (mongo.Session, error) {
	ss, err := wc.cc.StartSession(opts...)
	if err != nil {
		return nil, handle(err)
	}
	return &Session{Session: ss}, nil
}

func (wc *Client) UseSession(ctx context.Context, fn func(mongo.SessionContext) error) error {
	return handle(wc.cc.UseSession(ctx, fn))
}

func (wc *Client) UseSessionWithOptions(ctx context.Context, opts *options.SessionOptions, fn func(mongo.SessionContext) error) error {
	return handle(wc.cc.UseSessionWithOptions(ctx, opts, fn))
}

func (wc *Client) Client() *mongo.Client { return wc.cc.Client() }
