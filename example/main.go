package main

import (
	"fmt"
	"github.com/nafisfaysal/matterhook"
)

func main() {
	webhookURL := "https://matterhook.com/3423knkldv323"

	var msg matterhook.Message

	attachment := matterhook.Attachment{
		Text:  "Attact Text",
		Title: "Attact Title",
		Fields: []matterhook.Field{
			{
				Title: "Attach field Title",
			},
		},
	}

	anotherAttachment := matterhook.Attachment{
		Text:       "Hello",
		Title:      "No title",
		AuthorName: "nafis",
		AuthorIcon: "give author icon",
		Fields: []matterhook.Field{
			{
				Title: "Filed Title",
				Value: "Some value",
				Short: true,
			},
		},
	}

	msg.AddAttachments([]matterhook.Attachment{attachment, anotherAttachment})

	err := matterhook.Send(webhookURL, msg)
	if err != nil {
		fmt.Println(err)
	}
}
