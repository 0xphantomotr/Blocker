package node

import (
	"context"
	"fmt"

	"github.com/0xphantomotr/Blocker/proto"
	"google.golang.org/grpc/peer"
)

type Node struct {
	verson string
	proto.UnimplementedNodeServer
}

func NewNode() *Node {
	return &Node{
		verson: "blocker-0.1",
	}
}

func (n *Node) Handshake(ctx context.Context, v *proto.Version) (*proto.Version, error) {
	ourVersion := &proto.Version{
		Version: n.verson,
		Height:  100,
	}

	p, _ := peer.FromContext(ctx)

	fmt.Printf("received version from %s %+v\n", v, p.Addr)
	return ourVersion, nil
}

func (n *Node) HandleTransaction(ctx context.Context, tx *proto.Transaction) (*proto.Ack, error) {
	peer, _ := peer.FromContext(ctx)
	fmt.Println("received tx from:", peer)
	return &proto.Ack{}, nil
}
