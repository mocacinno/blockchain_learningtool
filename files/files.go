package files

import (
	"blockchain_learningtool/shared"
	"crypto/x509"
//	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/ssh"
	"log"
	"os"
	"crypto/rsa"
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

/*
func WriteIdentitysToFile(identities []shared.Identity) {
	fileName := "output/identities.txt"
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for _, id := range identities {
		privkeyfileName := "output/keys/" + id.Name 
		privkeyfile, err := os.Create(privkeyfileName)
		if err != nil {
			log.Fatal(err)
		}
		defer privkeyfile.Close()
		pubkeyfileName := "output/keys/" + id.Name + ".pub"
		pubkeyfile, err := os.Create(pubkeyfileName)
		if err != nil {
			log.Fatal(err)
		}
		defer pubkeyfile.Close()



		fmt.Fprintf(file, "\n\n\n\nName: %s\n", id.Name)

		pubKeyBlock := &pem.Block{
			Type:  "PUBLIC KEY",
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

		pubKeyArmored, err = armor.Encode(pubkeyfile, "PGP PUBLIC KEY BLOCK", nil) // Provide the correct armor type
		if err != nil {
			fmt.Printf("%s\n", err)
		}
		defer pubKeyArmored.Close()
		if err := pem.Encode(pubKeyArmored, pubKeyBlock); err != nil {
			log.Fatal(err)
		}


		//sshKey := fmt.Sprintf("ssh-rsa %s", base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PublicKey(id.PublicKey)))
		//fmt.Fprintln(file, "SSH Public Key:")
		//fmt.Fprintln(file, sshKey)

		//hexKey := fmt.Sprintf("0x%s", base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PublicKey(id.PublicKey)))
		//fmt.Fprintln(file, "Hexadecimal Public Key:")
		//fmt.Fprintln(file, hexKey)

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

		privKeyArmored, err = armor.Encode(privkeyfile, "PGP PRIVATE KEY BLOCK", nil) // Provide the correct armor type
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
*/
type KeyFormat int

const (
	PemFormat KeyFormat = iota
	SshFormat
	HexFormat
)

// WriteIdentitysToFile writes key information about each identity to files
func WriteIdentitysToFile(identities []shared.Identity, outputDir string) {
	//could change the format if i wanted to
	//formats := []KeyFormat{PemFormat, SshFormat, HexFormat}
	formats := []KeyFormat{PemFormat}
	for _, id := range identities {
		writePublicKey(id, outputDir, formats)
		writePrivateKey(id, outputDir, formats)
	}
}

func writePublicKey(id shared.Identity, outputDir string, formats []KeyFormat) {
	pubKeyFileName := fmt.Sprintf("%s/keys/%s.pub", outputDir, id.Name)
	pubKeyFile, err := os.Create(pubKeyFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer pubKeyFile.Close()

	pubKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(id.PublicKey),
	}

	for _, format := range formats {
		switch format {
		case PemFormat:
			writePEM(pubKeyFile, pubKeyBlock, "PGP PUBLIC KEY BLOCK")
		case SshFormat:
			writeSSH(pubKeyFile, id.PublicKey)
		case HexFormat:
			writeHex(pubKeyFile, id.PublicKey)
		}
	}
}

func writePrivateKey(id shared.Identity, outputDir string, formats []KeyFormat) {
	privKeyFileName := fmt.Sprintf("%s/keys/%s", outputDir, id.Name)
	privKeyFile, err := os.Create(privKeyFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer privKeyFile.Close()

	privKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(id.PrivateKey),
	}

	for _, format := range formats {
		switch format {
		case PemFormat:
			writePEM(privKeyFile, privKeyBlock, "PGP PRIVATE KEY BLOCK")
		case SshFormat:
			writeSSH(privKeyFile, id.PublicKey)
		case HexFormat:
			writeHex(privKeyFile, id.PublicKey)
		}
	}
}
/*
func writePEM(file *os.File, block *pem.Block, armorType string) {
	armoredFile, err := armor.Encode(file, armorType, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer armoredFile.Close()

	// Manually insert line breaks after every 64 characters
	data := pem.EncodeToMemory(block)

	// Print the content of data before encoding
	fmt.Println("Before encoding:", string(data))

	for len(data) > 0 {
		lineLength := 64
		if len(data) < lineLength {
			lineLength = len(data)
		}

		if _, err := armoredFile.Write(data[:lineLength]); err != nil {
			log.Fatal(err)
		}

		data = data[lineLength:]

		// Add a newline after each line, excluding the last line
		if len(data) > 0 {
			if _, err := armoredFile.Write([]byte("\n")); err != nil {
				log.Fatal(err)
			}
		}
	}
}

*/
func writePEM(file *os.File, block *pem.Block, armorType string) {
	armoredFile, err := armor.Encode(file, "PGP PRIVATE KEY BLOCK", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer armoredFile.Close()

	// Manually insert line breaks after every 64 characters
	data := pem.EncodeToMemory(block)
	fmt.Println("Before encoding:", string(data))

	lineLength := 64
	for i := 0; i < len(data); i += lineLength {
		end := i + lineLength
		if end > len(data) {
			end = len(data)
		}

		line := data[i:end]
		// Directly use armoredFile.Write without fmt.Fprintf
		if _, err := armoredFile.Write(line); err != nil {
			log.Fatal(err)
		}

		// Add newline after each line, excluding the last line
		if i+lineLength < len(data) {
			if _, err := armoredFile.Write([]byte("\n")); err != nil {
				log.Fatal(err)
			}
		}
	}
}




func writeSSH(file *os.File, pubKey *rsa.PublicKey) {
	// Convert the RSA public key to SSH key format
	sshKey, err := ssh.NewPublicKey(pubKey)
	if err != nil {
		log.Fatal(err)
	}

	// Write the SSH public key to the file
	sshKeyString := string(ssh.MarshalAuthorizedKey(sshKey)) + "\n"
	if _, err := file.WriteString(sshKeyString); err != nil {
		log.Fatal(err)
	}
}

func writeHex(file *os.File, pubKey *rsa.PublicKey) {
	// Convert the RSA public key to hex format
	hexKey := fmt.Sprintf("0x%x", pubKey.N)

	// Write the hex public key to the file
	if _, err := file.WriteString(hexKey + "\n"); err != nil {
		log.Fatal(err)
	}
}