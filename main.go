package main

import(
	"blockchain_learningtool/identity"
	"blockchain_learningtool/shared"
	"blockchain_learningtool/files"
)

func main() {
	myUsers := []string{"mocacinno", "bob", "alice", "james", "laura", "david", "emily", "alex", "sarah", "michael"}
	var userStruct []shared.Identity
	for _, currentuser := range(myUsers) {
		newidentityPtr := identity.GenerateIdentity(currentuser)
		newidentity := *newidentityPtr
		userStruct = append(userStruct, newidentity)
	}

	files.WriteIdentitysToFile(userStruct)
		
}
	
