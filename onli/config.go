package onli

const (
	// EnvDevelopment is the environment used by developers
	// to test and develop their applications. Usually is the more
	// unstable environment of all, because of constant changes and tests.
	EnvDevelopment = "development"

	// EnvStaging is the staging for confirmation and more business
	// focused tests. It may being shared with commercial and demos.
	EnvStaging = "staging"

	// EnvProduction is for production ready environment.
	EnvProduction = "production"

	// Defines the names of env vars to be used when loading configs,
	// applications can set these vars to configure the use of the sdk.
	audience     = "ONLI_AUDIENCE"
	clientId     = "ONLI_CLIENT_ID"
	clientSecret = "ONLI_CLIENT_SECRET"
	env          = "ONLI_ENVIRONMENT"
	scope        = "ONLI_SCOPE"
)

// config holds the parameters for configuration,
// see the documentation below to understand each param.
type config struct {
	// Audience param from RFC6749.
	Audience string

	// ClientID param from RFC6749.
	ClientID string

	// ClientSecret param from RFC6749.
	ClientSecret string

	// Env specifying the communication environment with the API.
	Env string

	// Scope param from RFC6749, a list of.
	Scope []string
}

// Config creates an empty config interface.
func Config() *config {
	return &config{}
}

// WithAudience sets the audience param into config.
func (cfg *config) WithAudience(audience string) *config {
	cfg.Audience = audience
	return cfg
}

// WithClientID sets the client id param into config.
func (cfg *config) WithClientID(clientId string) *config {
	cfg.ClientID = clientId
	return cfg
}

// WithClientSecret sets the client secret param into config.
func (cfg *config) WithClientSecret(clientSecret string) *config {
	cfg.ClientSecret = clientSecret
	return cfg
}

// WithEnv sets the env param into config.
func (cfg *config) WithEnv(env string) *config {
	cfg.Env = env
	return cfg
}

// WithScope sets the scope param into config.
func (cfg *config) WithScope(scope []string) *config {
	cfg.Scope = scope
	return cfg
}
