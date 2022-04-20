package server

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"port/pkg/grpc/pb"
)

func Start(config *ServiceConfig) error {
	log.Info("Starting GRPC service: " + config.url + ":" + config.port)
	listener, err := net.Listen("tcp", config.url+":"+config.port)
	if err != nil {
		log.Error("Can not listen on "+config.url+":"+config.port+", Error: ", err)
		return err
	}
	s := grpc.NewServer()
	pb.RegisterPortServiceServer(s, &service{dbHandler: config.dbHandler})

	if err := s.Serve(listener); err != nil {
		log.Error("listener Error: ", err)
		return err
	}
	return nil
}

func (e ErrNilHandler) Error() string {
	return "Database handler is nil"
}
