package service

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/selector"
	"github.com/micro/go-micro/server"
	k8s "github.com/micro/kubernetes/go/micro"

	gcli "github.com/micro/go-plugins/client/grpc"
	"github.com/micro/go-plugins/registry/noop"
	"github.com/micro/go-plugins/selector/named"
	"github.com/micro/go-plugins/selector/static"
	gsrv "github.com/micro/go-plugins/server/grpc"
)

// NewService returns a grpc service compatible with go-micro.Service
func NewService(address string, opts ...micro.Option) micro.Service {
	// our grpc client
	c := gcli.NewClient(client.Selector(static.NewSelector()))
	// our grpc server
	s := gsrv.NewServer(server.Address(address))

	// noop registry
	r := noop.NewRegistry()

	// named selector
	se := named.NewSelector(selector.Registry(r))

	// set selector
	c.Init(client.Selector(se))

	// create options with priority for our opts
	options := []micro.Option{
		micro.Client(c),
		micro.Server(s),
		micro.Registry(r),
	}

	// append passed in opts
	options = append(options, opts...)

	// generate and return a service
	return micro.NewService(options...)
}

func NewKubernetesService(address string, opts ...micro.Option) micro.Service {
	// our grpc client
	c := gcli.NewClient(client.Selector(static.NewSelector()))
	// our grpc server
	s := gsrv.NewServer(server.Address(address))
	s.Init()

	// noop registry
	r := noop.NewRegistry()

	// named selector
	se := named.NewSelector(selector.Registry(r))

	// set selector
	c.Init(client.Selector(se))

	// create options with priority for our opts
	options := []micro.Option{
		micro.Client(c),
		micro.Server(s),
		micro.Registry(r),
	}

	// append passed in opts
	options = append(options, opts...)

	// generate and return a service
	return k8s.NewService(options...)
}