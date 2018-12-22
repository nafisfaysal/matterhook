# Matterhook

Simple Go client for [Mattermost's WebHooks API](https://www.mattermost.org/webhooks/).

## Installation
```sh
$ go get -u https://github.com/nafisfaysal/matterhook
```

## Usage

Basic text message

```go
   webhookURL := "https://docs.mattermost.com/developer/webhooks-incoming.html"
   message := matterhook.Message{
      Text: "Hello MatterHook",
   }
   err := matterhook.Send(weebhokURL, message)
   if err != nil {
      fmt.Println(err)
   }

```

Send message with attachments
```go
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
```


Contributions are very welcome.
