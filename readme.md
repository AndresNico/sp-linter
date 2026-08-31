## How to use

- install `golangci-lint`: `go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.13.1`
- build a custom binary with the plugin baked in `golangci-lint custom --name sp-golangci-lint --destination ./bin`
- plugin linter needs to be enabled in .golangci.yaml
- Call the binary `./bin/sp-golangci-lint linters`. This should list only the `serviceprovider_linter` as enabled
