package files

import (
	"blockchain_learningtool/shared"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"golang.org/x/crypto/openpgp/armor"
	"log"
	"os"
)

func CreateDirs() error {
	toCreateDirs := []string{"output/keys", "output/blocks", "output/walktrough"}
	for _, directoryPath := range toCreateDirs {
		if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
			err := os.MkdirAll(directoryPath, 0755)
			if err != nil {
				return errors.New(fmt.Sprintf("Error creating directory: %s", err))
			}
			if shared.Myparameters.Verbose {
				fmt.Printf("Directory %s created successfully!\n", directoryPath)
			}

		} else {
			if shared.Myparameters.Verbose {
				fmt.Printf("Directory %s already exists.\n", directoryPath)
			}

		}
	}

	return nil
}

func WriteIdentitysToFile(identities []shared.Identity) {
	fileName := "output/identities.txt"
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for _, id := range identities {
		fmt.Fprintf(file, "\n\n\n\nName: %s\n", id.Name)

		pubKeyBlock := &pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(id.PublicKey),
		}
		pubKeyArmored, err := armor.Encode(file, "PGP PUBLIC KEY BLOCK", nil) // Provide the correct armor type
		if err != nil {
			fmt.Printf("%s\n", err)
		}
		defer pubKeyArmored.Close()
		if err := pem.Encode(pubKeyArmored, pubKeyBlock); err != nil {
			log.Fatal(err)
		}

		sshKey := fmt.Sprintf("ssh-rsa %s", base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PublicKey(id.PublicKey)))
		fmt.Fprintln(file, "SSH Public Key:")
		fmt.Fprintln(file, sshKey)

		hexKey := fmt.Sprintf("0x%s", base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PublicKey(id.PublicKey)))
		fmt.Fprintln(file, "Hexadecimal Public Key:")
		fmt.Fprintln(file, hexKey)

		privKeyBlock := &pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(id.PrivateKey),
		}
		privKeyArmored, err := armor.Encode(file, "PGP PRIVATE KEY BLOCK", nil) // Provide the correct armor type
		if err != nil {
			fmt.Printf("%s\n", err)
		}
		defer privKeyArmored.Close()
		if err := pem.Encode(privKeyArmored, privKeyBlock); err != nil {
			log.Fatal(err)
		}
	}
	if shared.Myparameters.Verbose {
		fmt.Printf("Identities written to %s\n", fileName)
	}

}
