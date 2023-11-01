package composers

import (
	helloworld "github.com/tuanpham197/test_repo"
	sctx "github.com/viettranx/service-context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"sendo/pkg/common"
)

func ComposeHelloRPCClient(serviceCtx sctx.ServiceContext) helloworld.GreeterClient {
	configComp := serviceCtx.MustGet("config").(common.Config)

	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.Dial(configComp.GetGRPCServerAddress(), opts)

	if err != nil {
		log.Fatal(err)
	}

	return helloworld.NewGreeterClient(clientConn)
}
