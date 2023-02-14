# Sensitive

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![NPM](https://img.shields.io/badge/Watsize-Library-289548)](https://github.com/buildingwatsize)
[![CodeQL](https://github.com/buildingwatsize/sensitive/actions/workflows/github-code-scanning/codeql/badge.svg?branch=main)](https://github.com/buildingwatsize/sensitive/actions/workflows/github-code-scanning/codeql)

Sensitive is a middleware for [GoFiber](https://gofiber.io/) to blind sensitive value like mobile no, citizen id, etc. by defined configuration. Useful for security policies.

## Table of Contents

- [Sensitive](#sensitive)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Versions](#versions)
    - [v0.1.0 - `2023-02-14`](#v010---2023-02-14)
  - [Signatures](#signatures)
  - [Examples](#examples)
  - [Config](#config)
  - [Default Config](#default-config)
  - [Example Usage](#example-usage)

## Installation

```bash
  go get -u github.com/buildingwatsize/sensitive
```

## Versions

### v0.1.0 - `2023-02-14`

- Blind text from `abcdefg` into `axxxxxg` (just show only the first and the last character with "x" as mark in the middle)
- Blinding by specific keys in response body
- Supported custom mark (default: "x")
- Debug Mode available via config `sensitive.New(sensitive.Config{ DebugMode: true })`

[...more](./CHANGELOG.md)

## Signatures

```go
func New(config ...Config) fiber.Handler
```

## Examples

```go
func main() {
  app := fiber.New()

  app.Use(sensitive.New(sensitive.Config{}))
  
  // ... Handlers ...
}
```

## Config

```go
// Config defines the config for middleware.
type Config struct {
  // Optional. Default: nil
  Next func(c *fiber.Ctx) bool

  // Required. Default: []
  Keys []string

  // Optional. Default: "x"
  Mark string

  // Optional. Default: false
  DebugMode bool
}
```

## Default Config

```go
var ConfigDefault = Config{
  Next:      nil,
  Keys:      []string{},
  Mark:      "x",
  DebugMode: false,
}
```

## Example Usage

Check it out! [example/README.md](./example/README.md)

---
made by ❤️ [buildingwatsize](https://github.com/buildingwatsize)
