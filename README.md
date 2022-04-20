# Port Service

This is Ports service which base on gRPC protocol.

### Development 

#### To regenerate protobuf / grpc stub
   ```bash
   make grpc
   ```

#### To build and push the image
   ```bash
    IMAGE=quay.io/arturobrzut/portservice:1.0 make docker-push
   ```

### Tests
#### to run unit test 
   ```bash
   make test
   ```
#### to run End to End test, first start the service 
   ```bash
   make run  
   make test-e2e
   ```

### Local run service
   ```bash
   make run
   ```
### Client application
- You can use any of grpc client app for Port service.
  - For example, you can use [grpc-client-cli](https://github.com/vadimi/grpc-client-cli)  to create, get or delete port record.
    
    ```bash
     grpc-client-cli --proto ./proto/port.proto :50051
    ```

### Docker Compose
Under development - use Docker compose to start container with pod service and with Mongo database.
- start port service and mongodb
  ```bash
  docker-compose up
  ```
- connect to service by client
  ```bash
  grpc-client-cli --proto ./proto/port.proto :50051
  ```


