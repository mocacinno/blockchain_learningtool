package identity

import (
	"blockchain_learningtool/shared"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

/*
func GenerateIdentity(name string) *shared.Identity {
	privKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	return &shared.Identity{
		Name:       name,
		PublicKey:  &privKey.PublicKey,
		PrivateKey: privKey,
	}
}
*/
func GenerateIdentity(name string) (*shared.Identity, error) {
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	// Save private key to file
	privKeyFile, err := os.Create("output/keys/" + name + "_private.pem")
	if err != nil {
		return nil, err
	}
	defer privKeyFile.Close()

	privKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privKey),
	}

	err = pem.Encode(privKeyFile, privKeyPEM)
	if err != nil {
		return nil, err
	}

	// Save public key to file
	pubKeyFile, err := os.Create("output/keys/" + name + "_public.pem")
	if err != nil {
		return nil, err
	}
	defer pubKeyFile.Close()

	pubKeyPEM := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&privKey.PublicKey),
	}

	err = pem.Encode(pubKeyFile, pubKeyPEM)
	if err != nil {
		return nil, err
	}

	return &shared.Identity{
		Name:       name,
		PublicKey:  &privKey.PublicKey,
		PrivateKey: privKey,
	}, nil
}
