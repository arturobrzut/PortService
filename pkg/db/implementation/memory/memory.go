package memory

import (
	log "github.com/sirupsen/logrus"
	"port/pkg/grpc/pb"
)

func New() (*Database, error) {
	setDbConfig()
	log.Info("Memory Database")
	return &Database{ports: make(map[string]*pb.Port)}, nil
}

func (*Database) Close() {
	log.Info("Close Database")
}

func (db *Database) Create(port *pb.Port) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	if _, ok := db.ports[port.Id]; ok {
		return ErrAlreadyExists{}
	}

	db.ports[port.Id] = port
	return nil
}

func (db *Database) Get(id string) (*pb.Port, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	port, ok := db.ports[id]
	if !ok {
		return nil, ErrNotFound{}
	}
	return port, nil
}

func (db *Database) Delete(id string) error {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	_, ok := db.ports[id]
	if !ok {
		return ErrNotFound{}
	}
	delete(db.ports, id)
	return nil
}

func (db *Database) Update(port *pb.Port) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	if _, ok := db.ports[port.Id]; !ok {
		return ErrNotFound{}
	}

	db.ports[port.Id] = port
	return nil
}

func (db *Database) CreateOrUpdate(port *pb.Port) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	db.ports[port.Id] = port
	return nil
}

func (e ErrAlreadyExists) Error() string {
	return "Already exists"
}

func (e ErrNotFound) Error() string {
	return "Not found"
}

func setDbConfig() map[int]string {
	dbConfig := map[int]string{}
	return dbConfig
}
