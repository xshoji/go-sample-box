package gateway

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/xshoji/go-sample-box/grpc-gateway/proto"
	"google.golang.org/grpc"
	"net/http"
)

func newGateway(ctx context.Context, grpEndpoint string, opts ...runtime.ServeMuxOption) (http.Handler, error) {
	mux := runtime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(grpEndpoint, dialOpts...)
	if err != nil {
		return nil, err
	}
	err = proto.RegisterUserServiceHandler(ctx, mux, conn)
	if err != nil {
		return nil, err
	}
	return mux, nil
}

// Run starts a HTTP server and blocks forever if successful.
func Run(address string, grpEndpoint string, opts ...runtime.ServeMuxOption) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gw, err := newGateway(ctx, grpEndpoint, opts...)
	if err != nil {
		return err
	}

	return http.ListenAndServe(address, gw)
}
