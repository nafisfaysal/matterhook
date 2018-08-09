package matterhook

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (m *Message) AddAttachment(attachment Attachment) *Message {
	m.Attachments = append(m.Attachments, attachment)
	return m
}

func (m *Message) AddAttachments(attachments []Attachment) *Message {
	m.Attachments = append(m.Attachments, attachments...)
	return m
}

func Send(url string, msg Message) error {
	payloadBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	resp.Body.Close()

	return nil
}
