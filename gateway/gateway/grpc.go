package gateway

import "github.com/shahriar-mohim007/commons/discovery"

type gateway struct {
	registry discovery.Registry
}

func NewGRPCGateway(registry discovery.Registry) *gateway {
	return &gateway{registry}
}
