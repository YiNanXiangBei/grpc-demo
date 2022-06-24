package register

import (
	"google.golang.org/grpc/resolver"
)

var (
//scheme      = "example"
//serviceName = "resolver.example.grpc.io"
//
//backendAddr = "localhost:50051"
)

type ResolverBuilder struct {
	scheme      string
	serviceName string
	backendAddr string
}

func NewResolverBuilder(scheme, serviceName, backendAddr string) *ResolverBuilder {
	return &ResolverBuilder{
		scheme:      scheme,
		serviceName: serviceName,
		backendAddr: backendAddr,
	}
}

func (b *ResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &grpcResolver{
		target: target,
		cc:     cc,
		addrsStore: map[string][]string{
			b.serviceName: {b.backendAddr},
		},
	}
	r.start()
	return r, nil
}

func (b *ResolverBuilder) Scheme() string {
	return b.scheme
}

type grpcResolver struct {
	target     resolver.Target
	cc         resolver.ClientConn
	addrsStore map[string][]string
}

func (g *grpcResolver) start() {
	addrStrs := g.addrsStore[g.target.URL.Host]
	addrs := make([]resolver.Address, len(addrStrs))
	for i, s := range addrStrs {
		addrs[i] = resolver.Address{Addr: s}
	}
	g.cc.UpdateState(resolver.State{Addresses: addrs})
}

func (*grpcResolver) ResolveNow(options resolver.ResolveNowOptions) {}

func (*grpcResolver) Close() {}

//func init() {
//	resolver.Register(&ResolverBuilder{})
//}
