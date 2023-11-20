package main

import (
	"fmt"
	"list_of_ingredients_producer/sap_api_wrapper"
	"list_of_ingredients_producer/teams_notifier"
	"list_of_ingredients_producer/utils"
	"time"
)

func main() {
	utils.LoadEnvs()

	fmt.Printf("%v: Started the Script pre-scheduler \n", time.Now().UTC().Format("2006-01-02 15:04:05"))
	err := utils.HandleAllItemsListOfIngredients()
	if err != nil {
		teams_notifier.SendUnknownErrorToTeams(err)
	}
	sap_api_wrapper.SapApiPostLogout()

	fmt.Printf("%v: Success \n", time.Now().UTC().Format("2006-01-02 15:04:05"))
}
