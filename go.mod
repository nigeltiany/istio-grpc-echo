module github.com/nigeltiany/micro_istio

replace github.com/nigeltiany/micro_istio v0.0.0 => ./

replace github.com/nigeltiany/micro-istio-internal v0.0.0 => ../micro-istio-internal

require github.com/nigeltiany/micro-istio-internal v0.0.0

require (
	github.com/golang/protobuf v1.2.0
	github.com/micro/cli v0.0.0-20180830071301-8b9d33ec2f19 // indirect
	github.com/micro/go-grpc v0.3.0 // indirect
	github.com/micro/go-micro v0.11.0
	github.com/micro/go-plugins v0.14.0
	github.com/micro/grpc-go v0.0.0-20180913204047-2c703400301b // indirect
	github.com/micro/kubernetes v0.1.0
	github.com/micro/mdns v0.0.0-20160929165650-cdf30746f9f7 // indirect
	golang.org/x/net v0.0.0-20180826012351-8a410e7b638d
	google.golang.org/grpc v1.15.0
)
