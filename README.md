# f1-fantasy-api-go

[![GoDoc](https://img.shields.io/badge/godoc-reference-green.svg)](https://godoc.org/github.com/vbonduro/f1-fantasy-api-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/vbonduro/f1-fantasy-api-go)](https://goreportcard.com/report/github.com/vbonduro/f1-fantasy-api-go)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/vbonduro/f1-fantasy-api-go/blob/main/LICENSE)

golang implementation of the [f1-fantasy-api](https://github.com/zeroclutch/f1-fantasy-api) project.

# Supported Features

The following APIs are currently supported:

* Get list of drivers & constructors
* Get list of circuits
* Retrieve leaderboard for a specific league

If you would like to see more features supported, please open an issue or send a pull request.

# Installation

`go get github.com/vbonduro/f1-fantasy-api-go`

# Usage

You can use either the public or authenticated APIs.

The public API can be instantiated with: `f1fantasy.NewApi()`. The Public API does not require login credentials but has
limited functionality.

The authenticated API can be instantiated with `f1fantasy.NewAuthenticatedApi(user,password)`. An authenticated API has access
to all functionality from this library.

All public API implementations can be found in the `pkg/f1fantasy` directory.

See [here](./test/main.go) for an example program.
