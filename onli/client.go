package onli

import (
	"context"
	"errors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"os"
	"strings"
)

type client struct {
	config       *config
	oauth2Client *clientcredentials.Config
	oauth2Token  *oauth2.Token
}

// Client initializes a new client enabling the use of services from
// the rest of sdk. It accepts a list of options to configure the client.
// It also loads by default, the configs from environment vars.
//
// The list of available env vars are the following:
//
// 	audience     = "ONLI_AUDIENCE"
//	clientId     = "ONLI_CLIENT_ID"
//	clientSecret = "ONLI_CLIENT_SECRET"
//	env          = "ONLI_ENVIRONMENT"
//	scope        = "ONLI_SCOPE"
//
// Some validations will occur, to check configuration validity,
// your application must check for the error that may return.
func Client(options ...*config) (*client, error) {
	c := new(client)
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

		if len(opt.Scope) != 0 {
			c.config.WithScope(opt.Scope)
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

	if len(c.config.Scope) == 0 {
		return nil, errors.New("no scope provided from config")
	}

	for _, s := range c.config.Scope {
		if s == "" {
			return nil, errors.New("scope contains empty values")
		}
	}

	return c, nil
}

// loadConfigFromEnv takes default env vars and load into config.
func (c *client) loadConfigFromEnv() {
	c.config.WithAudience(os.Getenv(audience))
	c.config.WithClientID(os.Getenv(clientId))
	c.config.WithClientSecret(os.Getenv(clientSecret))
	c.config.WithEnv(os.Getenv(env))
	c.config.WithScope(strings.Split(os.Getenv(scope), ","))
}

func (c client) Authorize(ctx context.Context) (accessToken string, err error) {
	// TODO
	return "", nil
}

func (c client) Ping(ctx context.Context) error {
	// TODO
	return nil
}
