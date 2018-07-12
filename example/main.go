package main

import (
	"fmt"
	matterhook "github.com/NafisFaysal/mattermost/mattermost"
)

func main() {
	weebhokURL := "https://matterhook.com/3423knkldv323"

	var attachment matterhook.Attachment
	attachment.Pretext = "Hello"
	attachment.Text = "No One"
	attachment.Title = "Test Title OIYOT"
	attachment.TitleLink = "https://www.google.com/"

	var msg matterhook.Message
	var AnotherAttachment matterhook.Attachment
	AnotherAttachment.Text = "New att"
	AnotherAttachment.Title = "Test Title WRERWREGWER"
	AnotherAttachment.TitleLink = "google.com"
	msg.AddAttachments([]matterhook.Attachment{attachment, AnotherAttachment})
	err := matterhook.Send(weebhokURL, msg)
	if err != nil {
		fmt.Println(err)
	}
}
