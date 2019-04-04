module grpctest

go 1.12

replace google.golang.org/grpc => github.com/yangjuncode/grpc-go v1.19.2

require (
	github.com/gogo/protobuf v1.2.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	golang.org/x/net v0.0.0-20190311183353-d8887717615a
	google.golang.org/grpc v0.0.0-00010101000000-000000000000
)
