package main

import (
	"fmt"
	matterhook "github.com/nafisfaysal/mattermost/mattermost"
)

func main() {
	weebhokURL := "https://matterhook.com/3423knkldv323"
        
	var msg matterhook.Message
	var attachment matterhook.Attachment
	attachment.Pretext = "Hello"
	attachment.Text = "No One"
	attachment.Title = "Test Title OIYOT"
	attachment.TitleLink = "https://www.google.com/"

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
