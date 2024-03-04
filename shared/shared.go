package shared

import (
	"crypto/rsa"
)

type Identity struct {
	Name       string
	PublicKey  *rsa.PublicKey
	PrivateKey *rsa.PrivateKey
}