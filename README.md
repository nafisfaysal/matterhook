# matterhook

## Features
* Send plain text and message with attachments


## Installation
```sh
$ go get -u https://github.com/nafisfaysal/matterhook
```

## Usage

Basic send text message

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
  webhookURL := "https://docs.mattermost.com/developer/webhooks-incoming.html"
  
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

```
