package mongodb

import (
	"context"

	tracewrap "github.com/opencensus-integrations/gomongowrapper"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(ctx context.Context, opts ...*options.ClientOptions) (*WrappedClient, error) {
	c, err := tracewrap.Connect(ctx, opts...)

	if err != nil {
		return nil, handle(err)
	}

	wc := &WrappedClient{cc: c}

	return wc, err
}
