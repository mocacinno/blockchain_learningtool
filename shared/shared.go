package shared

import (
	"crypto/rsa"
)

type Parameters struct {
	Verbose        bool
	Namelist       string
	InputValue     int
	NumberOfBlocks int
	Debug 			bool
}

var Myparameters Parameters

type Identity struct {
	Name           string
	PublicKey      *rsa.PublicKey  `json:"-"`
	PrivateKey     *rsa.PrivateKey `json:"-"`
	Unspentoutputs []UnspentOutput
	Id             int
}
/*
type Identity struct {
	Name           string
	PublicKey      *rsa.PublicKey  
	PrivateKey     *rsa.PrivateKey
	Unspentoutputs []UnspentOutput
	Id             int
}
*/
type UnspentOutput struct {
	Blocknumber int
	Linenumber  int
	Value       int
}
