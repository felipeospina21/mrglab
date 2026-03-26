package gitlab

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/felipeospina21/mrglab/internal/config"
	"github.com/hasura/go-graphql-client"
)

// Client wraps the GitLab GraphQL API client.
type Client struct {
	gql     *graphql.Client
	cfg     *config.Config
	devMode bool
}

// NewClient creates a new GitLab GraphQL client from the given config.
func NewClient(cfg *config.Config) *Client {
	if cfg.DevMode {
		return &Client{cfg: cfg, devMode: true}
	}

	httpClient := http.DefaultClient
	httpClient.Transport = &authedTransport{
		key:     cfg.APIToken,
		wrapped: http.DefaultTransport,
	}

	url := fmt.Sprintf("%s/api/graphql", cfg.BaseURL)
	return &Client{
		gql:     graphql.NewClient(url, httpClient),
		cfg:     cfg,
		devMode: false,
	}
}

func (c *Client) projectFullPath(projectID string) graphql.ID {
	idx := slices.IndexFunc(c.cfg.Filters.Projects, func(p config.Project) bool {
		return p.ID == projectID
	})
	return graphql.ID(c.cfg.Filters.Projects[idx].FullPath)
}

type authedTransport struct {
	key     string
	wrapped http.RoundTripper
}

func (t *authedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("PRIVATE-TOKEN", t.key)
	return t.wrapped.RoundTrip(req)
}
