package blocks

import (
	"blockchain_learningtool/shared"
	"blockchain_learningtool/signing"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func CreateBlockHeader(blocknumber int) []string {
	previousblocknumber := blocknumber - 1
	previousblockfilename := fmt.Sprintf("output/blocks/block%04d.csv", previousblocknumber)
	fileData, err := ioutil.ReadFile(previousblockfilename)
	if err != nil {
		log.Fatal(err)
	}

	hashSum := sha256.Sum256(fileData)
	output := []string{fmt.Sprintf("%x", hashSum), strconv.Itoa(blocknumber)}
	return output
}

func CreateNewBlock(blocknumber int, userStruct []shared.Identity) []shared.Identity {
	//first, the block header...
	filename := fmt.Sprintf("output/blocks/block%04d.csv", blocknumber)
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
	//now, let's create some random transactions, add them to the block, and update the userStruct by using UpdateUserAddUnspentoutputs and UpdateUserRemoveUnspentoutputs
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(5) + 1
	for transaction := 1; transaction <= randomNumber; transaction++ {
		UpdateduserStruct, txline := CreateNewTransaction(userStruct)
		userStruct = UpdateduserStruct
		if err := csvwriter.Write(txline); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
	if shared.Myparameters.Verbose {
		fmt.Printf("%s written\n", filename)
	}

	return userStruct
}

func CreateNewTransaction(userStruct []shared.Identity) ([]shared.Identity, []string) {
	if shared.Myparameters.Verbose {
		jstruct, _ := json.MarshalIndent(userStruct, "", "\t")
		fmt.Printf("we start with userstruct %s\n", jstruct)
	}

	var outputline []string
	//first, find a user with unspent outputs
	var nonEmptyEntries []shared.Identity
	for nonemptyIndex, id := range userStruct {
		if len(id.Unspentoutputs) > 0 {
			id.Id = nonemptyIndex
			nonEmptyEntries = append(nonEmptyEntries, id)
		}
	}
	if shared.Myparameters.Verbose {
		jstruct, _ := json.MarshalIndent(nonEmptyEntries, "", "\t")
		fmt.Printf("the non empty elements are %s\n", jstruct)
		fmt.Printf("Number of non-empty entries: %d\n", len(nonEmptyEntries))
	}

	rand.Seed(time.Now().UnixNano())
	randomIndexSelectedUser := rand.Intn(len(nonEmptyEntries))
	selectedEntry := nonEmptyEntries[randomIndexSelectedUser]
	if shared.Myparameters.Verbose {
		jstruct, _ := json.MarshalIndent(selectedEntry, "", "\t")
		fmt.Printf("the selected entry was %s\n", jstruct)
		fmt.Printf("user %s with index %d was selected for spending an unspent output\n", selectedEntry.Name, selectedEntry.Id)
	}

	//pick one, two or more of said unspent outputs (always pick one for the demo now)
	randomIndexUnspentOutput := rand.Intn(len(selectedEntry.Unspentoutputs))
	if shared.Myparameters.Verbose {
		fmt.Printf("unspent output index %d was selected from this user with id %d\n", randomIndexUnspentOutput, selectedEntry.Id)
	}

	//remove unspent output(s) from said user
	if shared.Myparameters.Verbose {
		fmt.Println("going to remove this output now")
	}

	UpdateduserStruct, value, blocknumber, linenumber, sender := UpdateUserRemoveUnspentoutputs(selectedEntry.Id, randomIndexUnspentOutput, userStruct)
	userStruct = UpdateduserStruct
	if shared.Myparameters.Verbose {
		jstruct, _ := json.MarshalIndent(userStruct, "", "\t")
		fmt.Printf("after removing the unspent output, the struct now is %s\n", jstruct)
	}

	//pick one or more receivers, split the value intput one or more parts (just transfer full value for the demo now)
	numberofreceivers := rand.Intn(5) + 1
	reveivervalues := make([]int, numberofreceivers)
	for i := 0; i < numberofreceivers; i++ {
		reveivervalues[i] = rand.Intn(value / numberofreceivers) // Random value for each user
	}
	totalSum := 0
	for _, v := range reveivervalues {
		totalSum += v
	}
	reveivervalues[numberofreceivers-1] += value - totalSum
	var outputs []string
	for receivernumber := 0; receivernumber < numberofreceivers; receivernumber++ {
		if reveivervalues[receivernumber] > 10 {
			if shared.Myparameters.Verbose {
				fmt.Printf("sending %d to user number %d\n", reveivervalues[receivernumber], receivernumber)
			}

			randomIndexReceiver := rand.Intn(len(userStruct))
			selectedEntry = userStruct[randomIndexReceiver]
			if shared.Myparameters.Verbose {
				fmt.Printf("user %s with id %d was selected as a receiver\n", selectedEntry.Name, randomIndexReceiver)
			}

			//use UpdateUserAddUnspentoutputs to add unspent output to receivers
			if shared.Myparameters.Verbose {
				fmt.Printf("adding unspent output to user %s (index %d), coming from block number %d, line number %d value %d\n", selectedEntry.Name, randomIndexReceiver, blocknumber, linenumber, reveivervalues[receivernumber])

			}
			UpdateduserStruct = UpdateUserAddUnspentoutputs(randomIndexReceiver, blocknumber, linenumber, reveivervalues[receivernumber], userStruct)
			userStruct = UpdateduserStruct
			if shared.Myparameters.Verbose {
				jstruct, _ := json.MarshalIndent(userStruct, "", "\t")
				fmt.Printf("after adding the unspent output, the struct now is %s\n", jstruct)
			}

			outputs = append(outputs, strconv.Itoa(reveivervalues[receivernumber]))
			outputs = append(outputs, selectedEntry.Name)
			outputs = append(outputs, base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PublicKey(selectedEntry.PublicKey)))
		}

	}

	//create transaction in block, put in []string and return
	outputline = append(outputline, "INPUTS")
	outputline = append(outputline, strconv.Itoa(blocknumber))
	outputline = append(outputline, strconv.Itoa(linenumber))
	outputline = append(outputline, strconv.Itoa(value))
	outputline = append(outputline, "SENDER")
	outputline = append(outputline, sender)
	outputline = append(outputline, "OUTPUTS")
	outputline = append(outputline, outputs...)
	outputline = append(outputline, "SIGNATURE")
	//outputline = append(outputline, "signature-todo")
	tosign := strings.Join(outputline, ",")
	signature, err := signing.SignMessage(selectedEntry.PrivateKey, tosign)
	if err != nil {
		fmt.Println("Error signing message:", err)
	}
	if shared.Myparameters.Verbose {
		fmt.Printf("as output, created tx csv line: %+v\n", outputline)
	}

	outputline = append(outputline, signature)
	return userStruct, outputline
}

func UpdateUserRemoveUnspentoutputs(IdentityIndex int, UnspentOutputIndex int, identities []shared.Identity) ([]shared.Identity, int, int, int, string) {
	unspentoutputslice := identities[IdentityIndex].Unspentoutputs
	value := unspentoutputslice[UnspentOutputIndex].Value
	blocknumber := unspentoutputslice[UnspentOutputIndex].Blocknumber
	linenumber := unspentoutputslice[UnspentOutputIndex].Linenumber
	sender := identities[IdentityIndex].Name
	if IdentityIndex < 0 || IdentityIndex >= len(identities) {
		fmt.Println("Invalid IdentityIndex")
		return identities, 0, 0, 0, ""
	}
	identity := identities[IdentityIndex]

	if UnspentOutputIndex < 0 || UnspentOutputIndex >= len(identity.Unspentoutputs) {
		fmt.Println("Invalid UnspentOutputIndex")
		return identities, 0, 0, 0, ""
	}

	identity.Unspentoutputs = append(identity.Unspentoutputs[:UnspentOutputIndex], identity.Unspentoutputs[UnspentOutputIndex+1:]...)

	identities[IdentityIndex] = identity

	return identities, value, blocknumber, linenumber, sender
}

func UpdateUserAddUnspentoutputs(indexnumber int, blocknumber int, linenumber int, value int, identities []shared.Identity) []shared.Identity {
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
func CreateInitialBlock(receiver shared.Identity, StartValue int) int {
	csvFile, err := os.Create("output/blocks/block0001.csv")
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
	row = []string{"INPUTS", "0", "0", strconv.Itoa(StartValue), "SENDER", "initialise", "OUTPUTS", strconv.Itoa(StartValue), receiver.Name, base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PublicKey(receiver.PublicKey)), "SIGNATURE", "initialise"}
	if err = csvwriter.Write(row); err != nil {
		log.Fatalln("error writing record to file", err)
	}
	if shared.Myparameters.Verbose {
		fmt.Println("output/blocks/block0001.csv written")
	}

	return StartValue
}
