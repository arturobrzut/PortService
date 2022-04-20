package e2e

import (
	"bufio"
	"context"
	"github.com/bcicen/jstream"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"os"
	"port/pkg/grpc/pb"
	"testing"
)

func TestUploadStatus(t *testing.T) {
	assert := assert.New(t)
	client, conn := startClient()
	defer conn.Close()
	response := LoadData(client)
	assert.Equal(true, response.Status, "Response should be true")
}

func TestCheckData(t *testing.T) {
	assert := assert.New(t)
	client, conn := startClient()
	defer conn.Close()
	name := CheckNameFromData(client, "VEBLA")
	assert.Equal("Barcelona", name, "Response should be Barcelona")
}

func startClient() (pb.PortServiceClient, *grpc.ClientConn) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Error("Client can not start", err)
		os.Exit(1)
	}

	return pb.NewPortServiceClient(conn), conn
}

func LoadData(grpcClient pb.PortServiceClient) *pb.Response {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	uploadClient, err := grpcClient.Upload(ctx)
	if err != nil {
		log.Error("Could not Upload data: ", err)
	}
	f, err := os.Open("ports.json")
	if err != nil {
		log.Fatalf("Error to read [file=%v]: %v", "port.json", err.Error())
	}

	r := bufio.NewReader(f)

	dec := jstream.NewDecoder(r, 1).EmitKV()

	for entry := range dec.Stream() {
		d := entry.Value.(jstream.KV)
		id, value := d.Key, d.Value.(map[string]interface{})

		port := pb.Port{}
		port.Id = id
		if value["name"] != nil {
			port.Name = value["name"].(string)
		}
		if value["city"] != nil {
			port.City = value["city"].(string)
		}
		if value["code"] != nil {
			port.Code = value["code"].(string)
		}
		err := uploadClient.Send(&port)
		if err != nil {
			log.Error("Could not send data: "+port.String()+", Error: ", err)
		} else {
			log.Info("Send data: " + port.String())
		}
	}

	response, err := uploadClient.CloseAndRecv()
	if err != nil {
		log.Error("could not close the sender: ", err)
	}
	log.Info("Response from server: ", response.String())
	return response

}

func CheckNameFromData(grpcClient pb.PortServiceClient, id string) string {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	port, err := grpcClient.Get(ctx, &pb.PortId{Id: id})
	if err != nil {
		log.Error("Could not get data: ", err)
	}

	return port.Name
}
