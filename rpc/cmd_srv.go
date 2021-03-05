package rpc

import (
	"context"
	"fmt"
	"github.com/otcChain/chord-go/p2p"
	"github.com/otcChain/chord-go/pbs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type cmdService struct{}

func newCmdRpc() *cmdService {
	return &cmdService{}
}

func (c *cmdService) P2PShowPeers(_ context.Context, peer *pbs.ShowPeer) (*pbs.CommonResponse, error) {
	network, ok := p2p.Inst().(*p2p.NetworkV1)
	if !ok {
		return nil, fmt.Errorf("this test case is not valaible")
	}
	result := network.DebugTopicPeers(peer.Topic)
	return &pbs.CommonResponse{
		Msg: result,
	}, nil
}

func (c *cmdService) P2PSendTopicMsg(_ context.Context, msg *pbs.TopicMsg) (*pbs.CommonResponse, error) {

	network, ok := p2p.Inst().(*p2p.NetworkV1)
	if !ok {
		return nil, fmt.Errorf("this test case is not valaible")
	}

	result := network.DebugTopicMsg(msg.Topic, msg.Msg)
	return &pbs.CommonResponse{
		Msg: result,
	}, nil
}

func (c *cmdService) StartRpc() error {
	var address = fmt.Sprintf("127.0.0.1:%d", config.CmdPort)
	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	cmdServer := grpc.NewServer()

	pbs.RegisterCmdServiceServer(cmdServer, c)

	reflection.Register(cmdServer)
	go func() {
		if err := cmdServer.Serve(l); err != nil {
			panic(err)
		}
	}()
	return nil
}

func DialToCmdService() pbs.CmdServiceClient {

	var address = fmt.Sprintf("127.0.0.1:%d", config.CmdPort)

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := pbs.NewCmdServiceClient(conn)

	return client
}
