package api

import (
	"net/http"

	"github.com/felipeospina21/mrglab/internal/config"
	"github.com/hasura/go-graphql-client"
)

type authedTransport struct {
	key     string
	wrapped http.RoundTripper
}

func (t *authedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("PRIVATE-TOKEN", t.key)
	return t.wrapped.RoundTrip(req)
}

func newClient() *graphql.Client {
	cfg := &config.GlobalConfig
	httpClient := http.DefaultClient
	httpClient.Transport = &authedTransport{
		key:     cfg.APIToken,
		wrapped: http.DefaultTransport,
	}

	url := buildURL(cfg)
	return graphql.NewClient(url, httpClient)
}
