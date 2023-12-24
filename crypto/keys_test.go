package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {
	privKey := GeneratePrivateKey()
	assert.Equal(t, len(privKey.Bytes()), PrivKeyLen)
	pubKey := privKey.Public()
	assert.Equal(t, len(pubKey.Bytes()), PubKeyLen)
}

func TestNewPrivateKeyFromString(t *testing.T) {
	// seed := make([]byte, 32)
	// io.ReadFull(rand.Reader, seed)
	// fmt.Println(hex.EncodeToString((seed)))
	var (
		seed       = "42eb304814b83a371ff9916e142c22003d8e69a48946096f2610dae65cfa03b9"
		privKey    = NewPrivateKeyFromString(seed)
		addressStr = "96b214eecd2a1a2e6f990c54fb3a085b34aa3e20"
	)
	pubKey := privKey.Public()
	address := pubKey.Address()
	assert.Equal(t, PrivKeyLen, len(privKey.Bytes()))
	assert.Equal(t, addressStr, address.String())
}

func TestPrivateKeySign(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.Public()
	msg := []byte("foo bar baz")

	sig := privKey.Sign(msg)
	valid := sig.Verify(pubKey, msg)

	assert.True(t, valid)

	// Test with invalid msg
	assert.False(t, sig.Verify(pubKey, []byte("foo")))

	// Test with invalid pubKey
	invalidPrivKey := GeneratePrivateKey()
	invalidPubKey := invalidPrivKey.Public()
	assert.False(t, sig.Verify(invalidPubKey, msg))

}

func TestPublicKeyToAddress(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.Public()
	address := pubKey.Address()

	assert.Equal(t, AddressLen, len(address.Bytes()))
	fmt.Println(address)
}
