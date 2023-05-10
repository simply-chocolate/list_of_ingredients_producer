package teams_notifier

import (
	"fmt"
	"os"

	goteamsnotify "github.com/atc0005/go-teams-notify/v2"
	"github.com/atc0005/go-teams-notify/v2/messagecard"
)

func SendRawMaterialNotFoundErrorToTeams(fatherItemCode string, ingredientItemCode string) {
	client := goteamsnotify.NewTeamsClient()
	webhook := os.Getenv("TEAMS_WEBHOOK_URL")

	card := messagecard.NewMessageCard()
	card.Title = "Raw Material Not Found"
	card.Text = fmt.Sprintf("Script failed to create a list of ingredients because one of the raw materials couldn't be found in rawMaterialsMap.<BR/>"+
		"**Father ItemCode**: %v<BR/>"+
		"**Ingredient ItemCode**: %v<BR/>"+fatherItemCode, ingredientItemCode)

	if err := client.Send(webhook, card); err != nil {
		fmt.Println("SendOrderErrorToTeams failed to send the error.")
	}

}
