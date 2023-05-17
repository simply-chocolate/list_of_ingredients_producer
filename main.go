package main

import (
	"fmt"
	"list_of_ingredients_producer/teams_notifier"
	"list_of_ingredients_producer/utils"
	"time"
)

func main() {
	utils.LoadEnvs()

	fmt.Printf("%v Started the Script \n", time.Now().UTC().Format("2006-01-02 15:04:05"))

	// TODO: The script should never be scheduled. It should be run manually.
	// TODO: We need to add an argument for HandleAllItemsListOfIngredients() to specify weather its a check run or if we need to update items in SAP
	err := utils.HandleAllItemsListOfIngredients()
	if err != nil {
		teams_notifier.SendUnknownErrorToTeams(err)
	}
	fmt.Printf("%v Success \n", time.Now().UTC().Format("2006-01-02 15:04:05"))

	/*
		fmt.Println("Started the Cron Scheduler at:" + time.Now().UTC().Format("2006-01-02 15:04:05"))
			s := gocron.NewScheduler(time.UTC)
			_, _ = s.Cron("0 5,10,15 * * *").SingletonMode().Do(func() {
				fmt.Printf("%v Started the Script \n", time.Now().UTC().Format("2006-01-02 15:04:05"))

				sap_api_wrapper.SapApiPostLogout()
				fmt.Printf("%v Success \n", time.Now().UTC().Format("2006-01-02 15:04:05"))
			})
			s.StartBlocking()
	*/
}
