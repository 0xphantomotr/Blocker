package types

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/0xphantomotr/Blocker/crypto"
	"github.com/0xphantomotr/Blocker/util"
	"github.com/stretchr/testify/assert"
)

func TestHashBlock(t *testing.T) {
	block := util.RandomBlock()
	hash := HashBlock(block)
	fmt.Println(hex.EncodeToString(hash))
	assert.Equal(t, 32, len(hash))
}

func TestSignBlock(t *testing.T) {
	var (
		block   = util.RandomBlock()
		privKey = crypto.GeneratePrivateKey()
		pubKey  = privKey.Public()
	)

	sig := SignBlock(privKey, block)
	assert.Equal(t, 64, len(sig.Bytes()))
	assert.True(t, sig.Verify(pubKey, HashBlock(block)))

}
