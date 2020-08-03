module github.com/kramanathan01/ship/ship-cli-consignment

go 1.14

require (
	github.com/kramanathan01/ship/ship-service-consignment v0.0.0-20200729213241-272853a4373b
	github.com/kramanathan01/ship/ship-service-vessel v0.0.0-20200803220720-29f75c1df345
	github.com/micro/go-micro/v2 v2.9.1
	google.golang.org/genproto v0.0.0-20200729003335-053ba62fc06f // indirect
	google.golang.org/grpc/examples v0.0.0-20200729200327-d6c4e49aab24 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
)

// replace github.com/kramanathan01/ship/ship-service-consignment => ../ship-service-consignment
