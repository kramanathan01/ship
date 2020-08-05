module github.com/kramanathan01/ship/ship-service-user

go 1.14

require (
	github.com/coreos/etcd v3.3.22+incompatible // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/golang/protobuf v1.4.2
	github.com/jmoiron/sqlx v1.2.0
	github.com/lib/pq v1.8.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/satori/go.uuid v1.2.0
	golang.org/x/crypto v0.0.0-20200728195943-123391ffb6de
	google.golang.org/protobuf v1.25.0
)

// replace github.com/kramanathan01/ship/ship-service-user => ../ship-service-user
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
