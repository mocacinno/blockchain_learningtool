package blocks
import (
	"blockchain_learningtool/shared"
    "encoding/csv"
    "log"
    "os"
    "strconv"
	"crypto/x509"
	"encoding/base64"
	"crypto/sha256"
    "fmt"
    "io/ioutil"
)

func CreateBlockHeader(blocknumber int) []string {
	previousblocknumber := blocknumber - 1
	previousblockfilename := fmt.Sprintf("output/block%04d.csv", previousblocknumber)
	fileData, err := ioutil.ReadFile(previousblockfilename)
    if err != nil {
        log.Fatal(err)
    }

    hashSum := sha256.Sum256(fileData)
	output  := []string{fmt.Sprintf("%x", hashSum), strconv.Itoa(blocknumber)}
	return output
}

func CreateNewBlock(blocknumber int, userStruct []shared.Identity) []shared.Identity {
	//first, the block header... 
	filename := fmt.Sprintf("output/block%04d.csv", blocknumber)
	blockheader := CreateBlockHeader(blocknumber)
	csvFile, err := os.Create(filename)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer csvFile.Close()
	csvwriter := csv.NewWriter(csvFile)
    defer csvwriter.Flush()
	if err := csvwriter.Write(blockheader); err != nil {
		log.Fatalln("error writing record to file", err)
	}
	return userStruct
}

func UpdateUserAddUnspentoutputs (indexnumber int, blocknumber int, linenumber int, value int, identities []shared.Identity) []shared.Identity {
	updateslice := identities[indexnumber]
	Unspentoutputslice := updateslice.Unspentoutputs
	var newUnspentOutput shared.UnspentOutput
	newUnspentOutput.Blocknumber = blocknumber
	newUnspentOutput.Linenumber = linenumber
	newUnspentOutput.Value = value
	Unspentoutputslice = append(Unspentoutputslice, newUnspentOutput)
	updateslice.Unspentoutputs = Unspentoutputslice
	identities[indexnumber] = updateslice
	return identities
}
func CreateInitialBlock(receiver shared.Identity) int{
	csvFile, err := os.Create("output/block0001.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer csvFile.Close()
	csvwriter := csv.NewWriter(csvFile)
    defer csvwriter.Flush()

	row := []string{"initialise", strconv.Itoa(0)}
        if err := csvwriter.Write(row); err != nil {
            log.Fatalln("error writing record to file", err)
        }
	row = []string{"INPUTS", "0","0","10000","SENDER", "initialise", "OUTPUTS", "10000",receiver.Name,base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PublicKey(receiver.PublicKey)), "SIGNATURE","initialise"}
        if err = csvwriter.Write(row); err != nil {
            log.Fatalln("error writing record to file", err)
        }
	fmt.Println("output/block0001.csv written\n")
	return 10000
}