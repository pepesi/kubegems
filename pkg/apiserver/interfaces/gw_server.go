package interfaces

import (
	"context"
	"flag"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	v1 "kubegems.io/kubegems/pkg/apiserver/interfaces/apis/v1"
)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
)

func RunHTTP() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := v1.RegisterTenantServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err != nil {
		panic(err)
	}
	if err := v1.RegisterClusterServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err != nil {
		panic(err)
	}

	if err := http.ListenAndServe(":8081", mux); err != nil {
		panic(err)
	}
}
