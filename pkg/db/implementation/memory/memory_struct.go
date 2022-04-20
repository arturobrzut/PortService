package memory

import (
	"port/pkg/grpc/pb"
	"sync"
)

type Database struct {
	ports map[string]*pb.Port
	mutex sync.RWMutex
}

type ErrNotFound struct{}

type ErrAlreadyExists struct{}
