package server

import (
	"context"
	log "github.com/sirupsen/logrus"
	"io"
	"port/pkg/grpc/pb"
)

func (s service) Create(ctx context.Context, p *pb.Port) (*pb.Response, error) {
	err := s.dbHandler.CreateOrUpdate(p)
	if err != nil {
		return &pb.Response{Status: false}, err
	}
	return &pb.Response{Status: true}, nil
}

func (s service) Update(ctx context.Context, p *pb.Port) (*pb.Response, error) {
	err := s.dbHandler.Update(p)
	if err != nil {
		return &pb.Response{Status: false}, err
	}
	return &pb.Response{Status: true}, nil
}

func (s *service) Upload(stream pb.PortService_UploadServer) error {
	log.Info("Start Upload")
	for {
		port, err := stream.Recv()
		if err == io.EOF {
			log.Info("Upload successfully")
			return stream.SendAndClose(&pb.Response{Status: true})
		}
		if err != nil {
			log.Error("Error during upload", err)
			return err
		}
		s.dbHandler.Create(port)
		if err != nil {
			log.Error("Error during upload", err)
			return err
		}
	}
}

func (s service) Delete(ctx context.Context, port *pb.Port) (*pb.Response, error) {
	log.Info("Delete")
	err := s.dbHandler.Delete(port.Id)
	if err != nil {
		log.Error("Error during delete, ", err)
		return nil, err
	}
	return &pb.Response{Status: true}, nil
}

func (s service) Get(ctx context.Context, id *pb.PortId) (*pb.Port, error) {
	log.Info("Get")
	port, err := s.dbHandler.Get(id.Id)
	if err != nil {
		log.Error("Error during Get, ", err)
		return nil, err
	}
	return port, nil

}

func (s service) Download(params *pb.Params, downloadServer pb.PortService_DownloadServer) error {

	panic("implement me")
}
