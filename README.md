# Go Wrapper for [Fortnite-API.com](https://fortnite-api.com)

[![Go Reference](https://pkg.go.dev/badge/github.com/BurakYs/go-fortnite-api.svg)](https://pkg.go.dev/github.com/BurakYs/go-fortnite-api)
[![Release](https://img.shields.io/github/v/release/BurakYs/go-fortnite-api)](https://github.com/BurakYs/go-fortnite-api/releases)
[![License](https://img.shields.io/github/license/BurakYs/go-fortnite-api)](LICENSE)
[![Discord](https://img.shields.io/discord/621452110558527502?label=Discord&logo=discord)](https://discord.gg/eysmvFT2rV)

A simple Go wrapper for [Fortnite-API.com](https://fortnite-api.com), with support for all available endpoints.

## Installation

```sh
go get github.com/BurakYs/go-fortnite-api
```

## Quick Example

```go
package main

import (
	"context"
	"fmt"

	"github.com/BurakYs/go-fortnite-api"
)

func main() {
	client := fortniteapi.NewClient(fortniteapi.LanguageEnglish, "your-api-key")

	flags := fortniteapi.CombineFlags(
		fortniteapi.FlagIncludePaths,
		fortniteapi.FlagIncludeGameplayTags,
	)

	searchParams := &fortniteapi.SearchBRCosmeticParams{
		Name:          "Peely",
		ResponseFlags: flags,
	}

	cosmetic, err := client.SearchBRCosmetic(context.TODO(), searchParams)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(cosmetic.ID)
}
```

## Links

- [API Documentation](https://dash.fortnite-api.com)
- [pkg.go.dev](https://pkg.go.dev/github.com/BurakYs/go-fortnite-api)
- [Discord](https://discord.gg/eysmvFT2rV)

## License

This project is licensed under the [MIT License](LICENSE).