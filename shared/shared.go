package shared

import (
	"crypto/rsa"
)

type Identity struct {
	Name       string
	PublicKey  *rsa.PublicKey `json:"-"`
	PrivateKey *rsa.PrivateKey `json:"-"`
	Unspentoutputs []UnspentOutput
	Id int
}

type UnspentOutput struct {
	Blocknumber int
	Linenumber int
	Value	   int
}