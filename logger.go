package mongodb

import (
	"emperror.dev/errors"
	"logur.dev/logur"
)

var errLog logur.Logger = logur.NoopLogger{}

func handle(err error, fields ...map[string]interface{}) error {
	if err != nil {
		errLog.Error(err.Error(), fields...)
	}

	return err
}

// SetLogger configures the global database logger.
func SetLogger(logger logur.Logger) error {
	if logger == nil {
		return errors.New("logger is nil")
	}

	errLog = logur.WithField(logger, "component", "mongodb")

	return nil
}
