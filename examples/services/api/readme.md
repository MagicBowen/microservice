## install gRPC and protobuffer for golang

### install protobuffer compiler
```sh
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.7.1/protobuf-cpp-3.7.1.tar.gz
tar zxvf protobuf-cpp-3.7.1.tar.gz
cd protobuf-cpp-3.7.1
./configure
make && sudo make install
```

### install go protobuffer generator

```sh
go get -u github.com/golang/protobuf/proto
go get -u github.com/golang/protobuf/protoc-gen-go
```

### instrall gRPC for goang
```sh
git clone https://github.com/grpc/grpc-go.git $GOPATH/src/google.golang.org/grpc 
git clone https://github.com/golang/net.git $GOPATH/src/golang.org/x/net    
git clone https://github.com/golang/text.git $GOPATH/src/golang.org/x/text    
git clone https://github.com/google/go-genproto.git $GOPATH/src/google.golang.org/genproto
cd $GOPATH/src/    
go install google.golang.org/grpc
```

### generate go rpc interface

```sh
protoc --go_out=plugins=grpc:. user.proto
```