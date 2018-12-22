package matterhook

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func catchErr(err error) {
	if err != nil {
		log.Fatalf("error occurred: %s", err)
	}
}

func serve(v *string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		catchErr(err)
		*v = string(b)
	}))
}

func TestSend(t *testing.T) {
	expt := ""
	srv := serve(&expt)

	tt := []struct {
		payload Message
		expt    string
	}{
		{
			payload: Message{
				Text: "hello matterhook",
			},

			expt: `{"text":"hello matterhook"}`,
		},
		{
			payload: Message{
				Text:      "Hello matterhook.\nAnother hello matterhook.",
				Username:  "nafis",
				IconEmoji: "none",
				Channel:   "thanksmessage",
			},

			expt: `{"text":"Hello matterhook.\nAnother hello matterhook.","username":"nafis","icon_emoji":"none","channel":"thanksmessage"}`,
		},
		{
			payload: Message{
				Text:      "Message text",
				Username:  "haname",
				IconEmoji: ":papa:",
				Channel:   "thanksmessage",
				Attachments: []Attachment{
					{
						Text:       "Attach Text",
						AuthorName: "nafisfaysal",
						AuthorLink: "https://www.github.com/nafisfaysal",
						Fields: []Field{
							{
								Title: "Field Title",
								Value: "I know",
								Short: true,
							},
						},
					},
				},
			},

			expt: `{"text":"Message text","username":"haname","icon_emoji":":papa:","channel":"thanksmessage","attachments":[{"text":"Attach Text","author_name":"nafisfaysal","author_link":"https://www.github.com/nafisfaysal","fields":[{"title":"Field Title","value":"I know","short":true}]}]}`,
		},
	}
	for i, tc := range tt {
		err := Send(srv.URL, tc.payload)
		catchErr(err)
		if expt != tc.expt {
			t.Fatalf("#%d: expected = %q, got %q", i+1, tc.expt, expt)
		}
	}
}
