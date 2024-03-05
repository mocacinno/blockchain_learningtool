package main

import (
	"blockchain_learningtool/blocks"
	"blockchain_learningtool/files"
	"blockchain_learningtool/flags"
	"blockchain_learningtool/identity"
	"blockchain_learningtool/shared"
	"fmt"
	"strings"
)

func main() {
	flags.ParseFlags()
	files.CreateDirs()
	myUsers := strings.Split(shared.Myparameters.Namelist, ",")
	var userStruct []shared.Identity
	for _, currentuser := range myUsers {
		newidentityPtr := identity.GenerateIdentity(currentuser)
		newidentity := *newidentityPtr
		userStruct = append(userStruct, newidentity)
	}

	files.WriteIdentitysToFile(userStruct)
	unspentoutputsvalue := blocks.CreateInitialBlock(userStruct[0], shared.Myparameters.InputValue)
	userStruct = blocks.UpdateUserAddUnspentoutputs(0, 0, 1, unspentoutputsvalue, userStruct)
	for blocknumber := 2; blocknumber <= shared.Myparameters.NumberOfBlocks; blocknumber++ {
		userStruct = blocks.CreateNewBlock(blocknumber, userStruct)
	}
	if shared.Myparameters.Verbose {
		fmt.Printf("%+v", userStruct)
	}

}
