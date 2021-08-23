package onli

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Client for Onli Seguros services.
type Client struct {
	// RestClient is not exported here to make an standard rest client
	// between the services. It holds some defaults to help
	// communicating with the API in a standard way. You should not
	// care about using it directly.
	restClient *resty.Client

	config       *config
	oauth2Config *clientcredentials.Config
	oauth2Token  *oauth2.Token
}

// NewClient initializes a client enabling the use of services from
// the rest of sdk. It accepts a list of options to configure the client.
// It also loads by default, the configs from environment vars.
//
// The list of available env vars are the following:
//
// 	audience     = "ONLI_AUDIENCE"
//	clientId     = "ONLI_CLIENT_ID"
//	clientSecret = "ONLI_CLIENT_SECRET"
//	env          = "ONLI_ENVIRONMENT"
//	scopes       = "ONLI_SCOPES"
//
// Some validations will occur, to check configuration validity,
// your application must check for the error that may return.
func NewClient(options ...*config) (*Client, error) {
	c := new(Client)
	c.config = new(config)

	// The default configuration will be loaded from env vars,
	// making the use of options just in exceptional cases.
	c.loadConfigFromEnv()

	// Range over options provided to replace configs.
	for _, opt := range options {
		if opt.Audience != "" {
			c.config.WithAudience(opt.Audience)
		}

		if opt.ClientID != "" {
			c.config.WithClientID(opt.ClientID)
		}

		if opt.ClientSecret != "" {
			c.config.WithClientSecret(opt.ClientSecret)
		}

		if opt.Env != "" {
			c.config.WithEnv(opt.Env)
		}

		if len(opt.Scopes) != 0 {
			c.config.WithScopes(opt.Scopes)
		}
	}

	if c.config.Audience == "" {
		return nil, errors.New("empty audience from config")
	}

	if c.config.ClientID == "" {
		return nil, errors.New("empty client id from config")
	}

	if c.config.ClientSecret == "" {
		return nil, errors.New("empty client secret from config")
	}

	if c.config.Env == "" {
		return nil, errors.New("empty env from config")
	}

	if c.config.Env != EnvDevelopment && c.config.Env != EnvStaging && c.config.Env != EnvProduction {
		return nil, errors.New("env not valid from config")
	}

	if len(c.config.Scopes) == 0 {
		return nil, errors.New("no scope provided from config")
	}

	for _, s := range c.config.Scopes {
		if s == "" {
			return nil, errors.New("scopes contains empty values")
		}
	}

	// Instantiate a oauth2 client with client credentials.
	v := url.Values{}
	v.Set("audience", c.config.Audience)
	c.oauth2Config = &clientcredentials.Config{
		ClientID:       c.config.ClientID,
		ClientSecret:   c.config.ClientSecret,
		Scopes:         c.config.Scopes,
		TokenURL:       c.tokenUrl(),
		EndpointParams: v,
		AuthStyle:      oauth2.AuthStyleInParams,
	}

	// Configure the default http rest client.
	c.restClient = resty.New().
		SetCloseConnection(true).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetHostURL(c.apiBaseUrl()).
		OnAfterResponse(responseMiddleware)

	return c, nil
}

// loadConfigFromEnv takes default env vars and load into config.
func (c *Client) loadConfigFromEnv() {
	c.config.WithAudience(os.Getenv(audience))
	c.config.WithClientID(os.Getenv(clientId))
	c.config.WithClientSecret(os.Getenv(clientSecret))
	c.config.WithEnv(os.Getenv(env))
	c.config.WithScopes(strings.Split(os.Getenv(scopes), ","))
}

// tokenUrl returns the token url for each environment.
func (c Client) tokenUrl() string {
	switch c.config.Env {
	case EnvDevelopment:
		return "https://auth-dev.onli.com.br/oauth2/token"
	case EnvStaging:
		return "https://auth-staging.onli.com.br/oauth2/token"
	case EnvProduction:
		return "https://auth.onli.com.br/oauth2/token"
	default:
		return ""
	}
}

// apiBaseUrl returns the api base url for each environment.
func (c Client) apiBaseUrl() string {
	switch c.config.Env {
	case EnvDevelopment:
		return "https://api-dev.onli.com.br"
	case EnvStaging:
		return "https://api-staging.onli.com.br"
	case EnvProduction:
		return "https://api.onli.com.br"
	default:
		return ""
	}
}

var (
	ErrUnauthorized        = errors.New("unauthorized")
	ErrNotFound            = errors.New("not found")
	ErrInternalServerError = errors.New("internal server error")
	ErrInvalidAPIResponse  = errors.New("invalid api response")
)

type errorResponse struct {
	Error string `json:"error"`
}

func (c *errorResponse) unmarshal(data []byte) {
	json.Unmarshal(data, c)
}

// responseMiddleware handles responses from Requests made in the services,
// allows to handle responses and errors in a normalized way.
func responseMiddleware(_ *resty.Client, r *resty.Response) error {
	switch r.StatusCode() {
	case http.StatusOK:
		return nil
	case http.StatusUnauthorized:
		return ErrUnauthorized
	case http.StatusNotFound:
		return ErrNotFound
	case http.StatusInternalServerError:
		return ErrInternalServerError
	case http.StatusBadRequest:
		var resp errorResponse
		resp.unmarshal(r.Body())
		return errors.New(resp.Error)
	default:
		return ErrInvalidAPIResponse
	}
}

// Authorize uses client credentials to retrieve a token.
// The provided context optionally controls which HTTP client is used.
// See the oauth2.HTTPClient variable.
func (c *Client) Authorize(ctx context.Context) (err error) {
	c.oauth2Token, err = c.oauth2Config.Token(ctx)
	return err
}

// Ping the source API to use as health check.
func (c *Client) Ping(ctx context.Context) error {
	err := c.Authorize(ctx)
	if err != nil {
		return err
	}

	// Makes a simple http request to health check endpoint.
	resp, err := c.Request(ctx).Get("/v1/health")
	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		return errors.New("not healthy")
	}

	return nil
}

// Request returns the API Rest Client to be used in services,
// you should not care about using it directly. It also handles
// the OAuth2 token refresh when applicable.
func (c *Client) Request(ctx context.Context) *resty.Request {
	if !c.oauth2Token.Valid() {
		c.Authorize(ctx)
	}

	return c.restClient.
		R().
		SetContext(ctx).
		SetAuthToken(c.oauth2Token.AccessToken)
}
