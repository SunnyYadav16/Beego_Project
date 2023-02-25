package initializers

import (
	"bee_project/utils"
	"fmt"
)

func OfficeConnector() {
	graphHelper := utils.NewGraphHelper()
	initializeGraph(graphHelper)

	user, userError := graphHelper.GetUser()
	utils.CheckError("Error Getting User! ", userError)
	fmt.Println(*user.GetDisplayName())

	token, tokenError := graphHelper.GetUserToken()
	utils.CheckError("Error Getting Token! ", tokenError)
	fmt.Println(*token)
}

func initializeGraph(graphHelper *utils.GraphHelper) {
	err := graphHelper.InitializeGraphForAppAuth()
	utils.CheckError("Error initializing Graph for app auth: ", err)
}
