package mongodb

import (
	tracewrap "github.com/opencensus-integrations/gomongowrapper"
)

func RegisterAllViews() error {
	return tracewrap.RegisterAllViews()
}

func UnregisterAllViews() {
	tracewrap.UnregisterAllViews()
}
