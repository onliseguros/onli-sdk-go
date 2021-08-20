# onli-sdk-go

Onli Seguros API Client enabling Go programs to interact with our services in a simple and uniform way.

## Coverage

This API client package covers most of the existing Onli Seguros API calls and is updated regularly
to add new and/or missing endpoints. Currently the following services are supported:

- [x] Applications
- [x] Award Emojis
- [x] Branches
- [ ] Discussions (threaded comments)

## Usage

```go
import "github.com/onliseguros/onli-sdk-go/onli"
```

Construct a new Onli Seguros client, then use the various services to
access different parts of the API:

```go
// Initializes a new client using env vars.
client, err := onli.Client()
if err != nil {
  log.Fatalf("Failed to create client: %v", err)
}
```

You can also construct a client using options: 

```go
client, err := onli.Client(
    onli.Config().WithAudience("yourAudience"),
    onli.Config().WithClientID("clientId"),
    onli.Config().WithClientSecret("clientSecret"),
    onli.Config().WithEnv(onli.EnvStaging),
    onli.Config().WithScope([]string{"scope1","scope2"}),
)
```

## Issues

- If you have an issue: report it on the [issue tracker](https://github.com/onliseguros/onli-sdk-go/issues)

## License

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at <http://www.apache.org/licenses/LICENSE-2.0>