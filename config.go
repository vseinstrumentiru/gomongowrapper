package gomongowrapper

import (
	"time"

	"emperror.dev/errors"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type ReadPref struct {
	Mode         readpref.Mode
	MaxStaleness *time.Duration
	Tags         []string
}

// Config holds information necessary for connecting to a database.
type Config struct {
	URI            string
	Hosts          []string
	ReplicaSet     *string
	User           string
	Pass           string
	Name           string
	ReadPreference *ReadPref

	Params map[string]string
}

// Validate checks that the configuration is valid.
func (c Config) Validate() error {
	if c.URI == "" && len(c.Hosts) == 0 {
		return errors.New("database hosts or uri is required")
	}

	if c.URI == "" && len(c.Hosts) > 0 {
		if c.User == "" {
			return errors.New("database user is required")
		}

		if c.Name == "" {
			return errors.New("database name is required")
		}
	}

	return nil
}
