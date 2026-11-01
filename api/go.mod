module github.com/Putnami/demo-observe-go/api

go 1.23

require (
	github.com/Putnami/demo-observe-go/pricing v0.0.0
	github.com/google/uuid v1.5.0
)

require golang.org/x/text v0.21.0 // indirect

replace github.com/Putnami/demo-observe-go/pricing => ../pricing
