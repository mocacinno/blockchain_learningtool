package files
import (
	"blockchain_learningtool/shared"
	"fmt"
	"log"
	"os"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
)


func WriteIdentitysToFile(identities []shared.Identity) {
	fileName := "output/identities.txt"
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write each identity to the file
	for _, id := range identities {
		fmt.Fprintf(file, "Name: %s\n", id.Name)

		// Convert public key to ASCII-armored format
		pubKeyBlock := &pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(id.PublicKey),
		}
		pubKeyArmored := armor.Encode(file, pubKeyBlock)
		defer pubKeyArmored.Close()
		fmt.Fprintln(pubKeyArmored, "")

		// Convert private key to ASCII-armored format
		privKeyBlock := &pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(id.PrivateKey),
		}
		privKeyArmored := armor.Encode(file, privKeyBlock)
		defer privKeyArmored.Close()
		fmt.Fprintln(privKeyArmored, "")
	}

}