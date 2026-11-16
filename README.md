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

The repository is also a Putnami workspace with declared specs:
`pricing/specs/order-total.json` and `api/specs/quote.json`, each requirement
bound to the Go test that proves it with `spectest.Proves`. CI runs
`putnami test,validate --all`, which writes the spec verification record, then
`putnami observe --upload --publish`, which sends the evidence, the
declarations a change touches and that record to Putnami. A pull request that
loses a proof that is verified on `main` shows it waiting for a person's
decision.

On `main`, the `publish` job builds the `api` binary and uploads it as an
artifact.

## Run locally

```bash
go test ./pricing/... ./api/...
putnami test,validate --all
ADDR=:8080 go run ./api
curl -s -X POST localhost:8080/quote -d '{"lines":[{"SKU":"a","UnitCents":500,"Quantity":2}]}'
```
