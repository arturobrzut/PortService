package db

import (
	"port/pkg/grpc/pb"
)

type Handler interface {
	Get(id string) (*pb.Port, error)
	Create(port *pb.Port) error
	CreateOrUpdate(port *pb.Port) error
	Update(port *pb.Port) error
	Delete(id string) error
	Close()
}

type ErrUnknownDbType struct{}
