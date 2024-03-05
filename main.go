package main

import(
	"blockchain_learningtool/identity"
	"blockchain_learningtool/shared"
	"blockchain_learningtool/files"
	"blockchain_learningtool/blocks"
	"blockchain_learningtool/flags"
	"fmt"
)

func main() {
	files.CreateDirs()
	myparameters := flags.ParseFlags()
	myUsers := []string{"mocacinno", "bob", "alice", "james", "laura", "david", "emily", "alex", "sarah", "michael"}
	var userStruct []shared.Identity
	for _, currentuser := range(myUsers) {
		newidentityPtr := identity.GenerateIdentity(currentuser)
		newidentity := *newidentityPtr
		userStruct = append(userStruct, newidentity)
	}

	files.WriteIdentitysToFile(userStruct)
	unspentoutputsvalue := blocks.CreateInitialBlock(userStruct[0],myparameters.InputValue)
	userStruct = blocks.UpdateUserAddUnspentoutputs(0, 0,1,unspentoutputsvalue, userStruct)
	for blocknumber := 2; blocknumber <= 10; blocknumber++ {
		userStruct = blocks.CreateNewBlock(blocknumber, userStruct)
	} 
	fmt.Printf("%+v", userStruct)
}
	
