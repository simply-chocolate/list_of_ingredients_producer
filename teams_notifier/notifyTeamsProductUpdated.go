package teams_notifier

import (
	"fmt"
	"os"

	goteamsnotify "github.com/atc0005/go-teams-notify/v2"
	"github.com/atc0005/go-teams-notify/v2/messagecard"
)

func SendProductUpdatedMessageToTemas(itemCode string, changedFields string) {
	client := goteamsnotify.NewTeamsClient()
	webhook := os.Getenv("TEAMS_WEBHOOK_URL")

	card := messagecard.NewMessageCard()
	card.Title = "Item Updated in SAP"
	card.Text = fmt.Sprintf("the script has updated the following item.<BR>"+
		"**ItemCode**: %v<BR>"+
		"**Changed fields**:<BR> %v ", itemCode, changedFields)

	if err := client.Send(webhook, card); err != nil {
		fmt.Println("SendOrderErrorToTeams failed to send the error.")
	}

}
