module github.com/gochenzl/chess

go 1.23.1

require (
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/satori/go.uuid v1.2.0
	github.com/syndtr/goleveldb v1.0.0
	google.golang.org/protobuf v1.34.2
)

require (
	github.com/golang/snappy v0.0.0-20180518054509-2e65f85255db // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.34.2 // indirect
)

replace github.com/gochenzl/chess => ./chess

replace github.com/gochenzl => ../gochenzl
