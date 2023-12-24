package main

import (
	"context"
	"log"
	"time"

	"github.com/0xphantomotr/Blocker/crypto"
	"github.com/0xphantomotr/Blocker/node"
	"github.com/0xphantomotr/Blocker/proto"
	"github.com/0xphantomotr/Blocker/util"
	"google.golang.org/grpc"
)

func main() {
	makeNode(":3000", []string{}, true)
	time.Sleep(time.Second)
	makeNode(":4000", []string{":3000"}, false)
	time.Sleep(time.Second)
	makeNode(":5000", []string{":4000"}, false)

	for {
		time.Sleep(time.Millisecond * 100)
		makeTransaction()
	}

	select {}
}

func makeNode(listenAddr string, bootstrapNodes []string, isValidator bool) *node.Node {
	cfg := node.ServerConfig{
		Version:    "Blocker-1",
		ListenAddr: listenAddr,
	}
	if isValidator {
		cfg.PrivateKey = crypto.GeneratePrivateKey()
	}
	n := node.NewNode(cfg)
	go n.Start(listenAddr, bootstrapNodes)
	return n
}

func makeTransaction() {
	client, err := grpc.Dial(":3000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := proto.NewNodeClient(client)
	privKey := crypto.GeneratePrivateKey()

	tx := &proto.Transaction{
		Version: 1,
		Inputs: []*proto.TxInput{
			{
				PrevTxHash:   util.RandomHash(),
				PrevOutIndex: 0,
				PublicKey:    privKey.Public().Bytes(),
			},
		},
		Outputs: []*proto.TxOutput{
			{
				Amount:  99,
				Address: privKey.Public().Address().Bytes(),
			},
		},
	}
	// version := &proto.Version{
	// 	Version:    "blocker-0.1",
	// 	Height:     1,
	// 	ListenAddr: ":4000",
	// }

	_, err = c.HandleTransaction(context.TODO(), tx)
	if err != nil {
		log.Fatal((err))
	}
}
