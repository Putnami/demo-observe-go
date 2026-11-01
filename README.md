# demo-observe-go

A Go multi-module repository (`go.work`) that Putnami observes from the outside.
It keeps its own CI on GitHub Actions and emits standard reports.

| Module | Path | Depends on |
| --- | --- | --- |
| `pricing` | `pricing/` | `golang.org/x/text` |
| `api` | `api/` | `pricing`, `github.com/google/uuid` |

## CI

`.github/workflows/ci.yml` builds both modules, runs the tests with
`gotestsum`, and writes:

- `reports/junit.xml`: JUnit test results;
- `reports/coverage.out`: the Go coverage profile.

On `main`, the `publish` job builds the `api` binary and uploads it as an
artifact.

## Run locally

```bash
go test ./pricing/... ./api/...
ADDR=:8080 go run ./api
curl -s -X POST localhost:8080/quote -d '{"lines":[{"SKU":"a","UnitCents":500,"Quantity":2}]}'
```
