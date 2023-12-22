package types

import (
	"crypto/sha256"

	"github.com/0xphantomotr/Blocker/crypto"
	"github.com/0xphantomotr/Blocker/proto"

	pb "github.com/golang/protobuf/proto"
)

func SignBlock(pk *crypto.PrivateKey, b *proto.Block) *crypto.Signature {
	return pk.Sign(HashBlock(b))
}

// returns SHA256 of the header
func HashBlock(block *proto.Block) []byte {
	return HashHeader(block.Header)
}

func HashHeader(header *proto.Header) []byte {
	b, err := pb.Marshal(header)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(b)
	return hash[:]
}
