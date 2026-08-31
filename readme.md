## How to use

- Install `golangci-lint`: `go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.13.1`
- Build a custom binary with the plugin baked in `golangci-lint custom --name sp-golangci-lint --destination ./example/bin`
- Call the binary `./example/bin/sp-golangci-lint linters`. This should list only the `serviceprovider_linter` as enabled
- Change into the example directory `cd example`and run the linter `./bin/sp-golangci-lint run ./...`. This should report one issue on the example/bad/main.go file.