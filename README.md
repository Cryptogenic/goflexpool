# goflexpool
[![API Version Compatibility](https://img.shields.io/badge/API%20Version%20Compatibility-v1.0-green.svg)]()

goflexpool is an unofficial Golang binding for interacting with the [flexpool](https://flexpool.io) API. It provides wrappers for accessing all endpoints that are available in Flexpool v1.5's API as well as structures and helper utility functions. There are also toy examples that demonstrate how this library can be used and the type of information that can be extracted. These examples include:

`poolinfo` - Gets information about the pool including hashrate info, PPLNS share window, uncle rate, and average blocks per day.

`minerinfo` - Gets information about a miner on the pool, including their meta-details, worker information, payments, and blocks mined.

To get a full listing, see the generated [godocs](#). To see more information about usage of the example binaries, see their respective readme files.

## Getting Started
This project has cross-platform build support out of the box due to the nature of Golang. To use this library in your own project, simply use `go get` to install and/or update this library.

```
go get -u github.com/cryptogenic/goflexpool/pkg/api
```

## Packages
### api
The `api` package includes all the relevant structures and wrapper functions for interacting directly with the flexpool API. The structures/schema has remained true to the API with one exception: wei have been implicitly converted to gwei to avoid overflows and make it easier when calculating block rewards.

All endpoints that return a balance or involve a currency-related value should be assumed to be in gwei units unless otherwise specified in the documentation.

Similarly, all endpoints that return hashrate data are in hashes/second. Both currency and hashrates can be converted using the utils package which is also included in this repo.

### utils
The `utils` package includes helpful functions for converting currency and hashrates, as well as pool-related calcuation functions. This package might be expanded upon as time goes on.

## License
This project is licensed under the MIT license - see the [LICENSE](LICENSE.md) file for details.
