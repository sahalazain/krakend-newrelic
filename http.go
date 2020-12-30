package metrics

import (
	"context"
	"net/http"

	"github.com/devopsfaith/krakend/transport/http/client"
	newrelic "github.com/newrelic/go-agent"
)

// HTTPClientFactory includes a http.RoundTripper for NewRelic instrumentation
func HTTPClientFactory(cf client.HTTPClientFactory) client.HTTPClientFactory {
	return func(ctx context.Context) *http.Client {
		client := cf(ctx)

		if tx, ok := ctx.Value(nrCtxKey).(newrelic.Transaction); ok {
			client.Transport = newrelic.NewRoundTripper(tx, client.Transport)
		}

		return client
	}
}
