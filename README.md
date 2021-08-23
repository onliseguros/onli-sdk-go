# onli-sdk-go

[![GitHubAction Status](https://github.com/onliseguros/onli-sdk-go/workflows/test/badge.svg?branch=master)](https://github.com/onliseguros/onli-sdk-go/actions)

---

Onli Seguros API Client enabling Go programs to interact with our services in a simple and uniform way.

## Coverage

This API client package covers most of the existing Onli Seguros API calls and is updated regularly
to add new and/or missing endpoints. Currently the following services are supported:

### Brokers

- [x] Merchant
- [x] Channel
- [x] Product
- [ ] Lead
- [ ] Quote
- [ ] Sale
- [ ] Charge

### Customers

- [ ] Recommendation

### Insurers

- [ ] Effect
- [ ] Proposal
- [ ] Claim
- [ ] Policy

## Usage

```go
import "github.com/onliseguros/onli-sdk-go/onli"
```

Initialize a new Onli Seguros client using default environment variables, then use the various services to
access different parts of the API:

```go
client, err := onli.NewClient()
if err != nil {
    panic(err)
}
```

You can also construct a client using options: 

```go
client, err := onli.NewClient(
    onli.Config().WithAudience("yourAudience"),
    onli.Config().WithClientID("clientId"),
    onli.Config().WithClientSecret("clientSecret"),
    onli.Config().WithEnv(onli.EnvStaging),
    onli.Config().WithScopes([]string{"scope1","scope2"}),
)
```

Here is an example of how to use services with the above created client:

```go
import "github.com/onliseguros/onli-sdk-go/service/merchant"
```

```go
svc := merchant.New(client)

resp, err := svc.GetStore(ctx, "yourBrokerChannelId")
if err != nil {
    panic(err)
}
```

## Issues

If you have an issue: report it on the [issue tracker](https://github.com/onliseguros/onli-sdk-go/issues)

## License

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance 
with the License. You may obtain a copy of the License at <http://www.apache.org/licenses/LICENSE-2.0>
