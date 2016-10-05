package nrh

import (
	"net/http"

	"github.com/newrelic/go-agent"
)

type Options struct {
	ApplicationName string
	LicenseKey      string
}

type NewRelic struct {
	app *newrelic.Application
}

func New(options Options) (*NewRelic, error) {
	var app newrelic.Application
	app, err := newrelic.NewApplication(
		newrelic.NewConfig(
			options.ApplicationName,
			options.LicenseKey,
		),
	)
	if err != nil {
		return nil, err
	}
	return &NewRelic{&app}, nil
}

func (nr *NewRelic) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		txn := (*nr.app).StartTransaction(r.URL.Path, w, r)
		defer txn.End()
		h.ServeHTTP(w, r)
	})
}
