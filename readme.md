## How to use

Config
- .custom-gcl.yaml -> Defines plugin and where it comes from (github, local)
- call `golangci-lint custom` -> builds custom binary with plugin baked in
- plugin linter needs to be enabled in .golanci.yaml
- Call the binary `./custom-glc linters`