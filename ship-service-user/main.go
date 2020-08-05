package main

import (
	"log"

	ub "github.com/kramanathan01/ship/ship-service-user/proto/user"
	"github.com/micro/go-micro/v2"
)

const schema = `
	create table if not exists users (
		id varchar(36) not null,
		name varchar(125) not null,
		email varchar(225) not null unique,
		password varchar(225) not null,
		company varchar(125),
		primary key (id)
	);
`

func main() {

	db, err := NewConnection()
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	db.MustExec(schema)

	repo := NewPostgresRepository(db)

	service := micro.NewService(
		micro.Name("ship.service.user"),
	)

	service.Init()
	if err := ub.RegisterUserServiceHandler(service.Server(), &handler{repo}); err != nil {
		log.Panic(err)
	}

	if err := service.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
