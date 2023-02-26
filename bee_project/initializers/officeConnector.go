package initializers

import (
	"bee_project/utils"
	"fmt"
)

func OfficeConnector() string {
	graphHelper := utils.NewGraphHelper()
	initializeGraph(graphHelper)

	user, userError := graphHelper.GetUser()
	if userError != nil {
		utils.CheckError("Error Getting User! ", userError)
	}
	fmt.Println(*user.GetDisplayName())

	token, tokenError := graphHelper.GetUserToken()
	if tokenError != nil {
		utils.CheckError("Error Getting Token! ", tokenError)
	}
	//fmt.Println(*token)

	return *token
}

func initializeGraph(graphHelper *utils.GraphHelper) {
	err := graphHelper.InitializeGraphForAppAuth()
	utils.CheckError("Error initializing Graph for app auth: ", err)
}
