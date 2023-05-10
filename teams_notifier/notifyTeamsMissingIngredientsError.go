package teams_notifier

import (
	"fmt"
	"os"

	goteamsnotify "github.com/atc0005/go-teams-notify/v2"
	"github.com/atc0005/go-teams-notify/v2/messagecard"
)

func SendMissingIngredientsToTeams(fatherItemCode string, ingredientItemCode string, languageCode string) {
	client := goteamsnotify.NewTeamsClient()
	webhook := os.Getenv("TEAMS_WEBHOOK_URL")

	card := messagecard.NewMessageCard()
	card.Title = "Missing Ingredients"
	card.Text = fmt.Sprintf("Script failed to create a list of ingredients because one of the ingredients used didn't have a list of ingredients itself.<BR/>"+
		"**Father ItemCode**: %v<BR/>"+
		"**Ingredient ItemCode**: %v<BR/>"+
		"**LanguageCode**:%v<BR/>", fatherItemCode, ingredientItemCode, languageCode)

	if err := client.Send(webhook, card); err != nil {
		fmt.Println("SendOrderErrorToTeams failed to send the error.")
	}

}
