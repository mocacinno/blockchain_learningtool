package identity

import(
	"blockchain_learningtool/shared"
	"crypto/rand"
	"crypto/rsa"
	//"crypto/sha256"
	//"encoding/hex"
)

func GenerateIdentity(name string) *shared.Identity {
	privKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	return &shared.Identity{
		Name:       name,
		PublicKey:  &privKey.PublicKey,
		PrivateKey: privKey,
	}
}