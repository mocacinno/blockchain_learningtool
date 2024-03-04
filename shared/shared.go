package shared

import (
	"crypto/rsa"
)

type Identity struct {
	Name       string
	PublicKey  *rsa.PublicKey
	PrivateKey *rsa.PrivateKey
	Unspentoutputs []UnspentOutput
}

type UnspentOutput struct {
	Blocknumber int
	Linenumber int
	Value	   int
}