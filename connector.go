package gomongowrapper

import (
	"strings"

	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/tag"
)

// NewConnector returns a new database connector for the application.
func NewConnector(config Config) (*Client, error) {
	opts := options.Client()

	if config.URI != "" {
		opts = opts.ApplyURI(config.URI)
	}

	if len(config.Hosts) > 0 {
		opts = opts.SetHosts(config.Hosts)
	}

	if config.User != "" {
		opts = opts.SetAuth(options.Credential{
			AuthSource: config.Name,
			Username:   config.User,
			Password:   config.Pass,
		})
	}

	if config.ReplicaSet != nil {
		opts.ReplicaSet = config.ReplicaSet
	}

	if config.ReadPreference != nil {
		var rpOpts []readpref.Option

		if config.ReadPreference.MaxStaleness != nil {
			rpOpts = append(rpOpts, readpref.WithMaxStaleness(*config.ReadPreference.MaxStaleness))
		}

		if len(config.ReadPreference.Tags) > 0 {
			var tagSet tag.Set
			for _, t := range config.ReadPreference.Tags {
				if t == "" {
					tagSet = append(tagSet, tag.Tag{})
					break
				}
				kv := strings.Split(t, ":")
				if len(kv) != 2 {
					continue
				}
				tagSet = append(tagSet, tag.Tag{Name: kv[0], Value: kv[1]})
			}

			if len(tagSet) > 0 {
				rpOpts = append(rpOpts, readpref.WithTagSets(tagSet))
			}
		}

		var err error
		opts.ReadPreference, err = readpref.New(config.ReadPreference.Mode, rpOpts...)
		if err != nil {
			return nil, err
		}
	}

	return NewClient(opts)
}
