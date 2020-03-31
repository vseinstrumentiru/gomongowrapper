package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewConnector returns a new database connector for the application.
func NewConnector(config Config) (*WrappedClient, error) {
	opts := options.Client().
		SetHosts(config.Hosts).
		SetAuth(options.Credential{
			AuthSource: config.Name,
			Username:   config.User,
			Password:   config.Pass,
		})

	if config.ReplicaSet != nil {
		opts.ReplicaSet = config.ReplicaSet
	}

	return NewClient(opts)
}
