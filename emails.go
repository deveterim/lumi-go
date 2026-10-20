package lumi

import (
	"context"
	"strings"
)

type Tag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Attachment struct {
	Content     []byte `json:"content,omitempty"`
	Filename    string `json:"filename"`
	Path        string `json:"path,omitempty"`
	ContentType string `json:"content_type,omitempty"`
}

type SendEmailRequest struct {
	From        string            `json:"from"`
	To          []string          `json:"to"`
	Subject     string            `json:"subject"`
	Html        string            `json:"html,omitempty"`
	Text        string            `json:"text,omitempty"`
	Cc          []string          `json:"cc,omitempty"`
	Bcc         []string          `json:"bcc,omitempty"`
	ReplyTo     string            `json:"reply_to,omitempty"`
	Headers     map[string]string `json:"headers,omitempty"`
	Attachments []Attachment      `json:"attachments,omitempty"`
	Tags        []Tag             `json:"tags,omitempty"`
	ScheduledAt string            `json:"scheduled_at,omitempty"`
}

type SendEmailResponse struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

type Email struct {
	Id        string `json:"id"`
	From      string `json:"from"`
	To        any    `json:"to"`
	Subject   string `json:"subject"`
	Html      string `json:"html,omitempty"`
	Text      string `json:"text,omitempty"`
	Status    string `json:"status"`
	CreatedAt string `json:"createdAt"`
	LastEvent string `json:"last_event,omitempty"`
}

type EmailsSvc struct {
	client *Client
}

func (s *EmailsSvc) Send(params *SendEmailRequest) (*SendEmailResponse, error) {
	return s.SendWithContext(context.Background(), params)
}

func (s *EmailsSvc) SendWithContext(ctx context.Context, params *SendEmailRequest) (*SendEmailResponse, error) {
	payload := map[string]any{
		"from":    params.From,
		"subject": params.Subject,
	}
	if len(params.To) == 1 {
		payload["to"] = params.To[0]
	} else if len(params.To) > 1 {
		payload["to"] = params.To
	}
	if params.Html != "" {
		payload["html"] = params.Html
	}
	if params.Text != "" {
		payload["text"] = params.Text
	}
	if params.ReplyTo != "" {
		payload["reply_to"] = params.ReplyTo
	}
	if len(params.Cc) > 0 {
		payload["cc"] = params.Cc
	}
	if len(params.Bcc) > 0 {
		payload["bcc"] = params.Bcc
	}
	if len(params.Headers) > 0 {
		payload["headers"] = params.Headers
	}
	if len(params.Attachments) > 0 {
		payload["attachments"] = params.Attachments
	}
	if len(params.Tags) > 0 {
		payload["tags"] = params.Tags
	}
	if params.ScheduledAt != "" {
		payload["scheduled_at"] = params.ScheduledAt
	}

	sendPath := "/api/v1/send"
	if strings.HasSuffix(s.client.baseURL, "/v1") {
		sendPath = "/send"
	}
	req, err := s.client.NewRequest(ctx, "POST", sendPath, payload)
	if err != nil {
		return nil, err
	}

	var res struct {
		Data   SendEmailResponse `json:"data"`
		Id     string            `json:"id"`
		Status string            `json:"status"`
	}
	if err := s.client.Perform(req, &res); err != nil {
		// Fallback to /emails
		if req2, err2 := s.client.NewRequest(ctx, "POST", "/emails", payload); err2 == nil {
			if err3 := s.client.Perform(req2, &res); err3 == nil {
				err = nil
			}
		}
		if err != nil {
			return nil, err
		}
	}
	if res.Data.Id != "" {
		return &res.Data, nil
	}
	return &SendEmailResponse{Id: res.Id, Status: res.Status}, nil
}

func (s *EmailsSvc) Get(emailId string) (*Email, error) {
	return s.GetWithContext(context.Background(), emailId)
}

func (s *EmailsSvc) GetWithContext(ctx context.Context, emailId string) (*Email, error) {
	getPath := "/api/v1/send/" + emailId
	if strings.HasSuffix(s.client.baseURL, "/v1") {
		getPath = "/send/" + emailId
	}
	req, err := s.client.NewRequest(ctx, "GET", getPath, nil)
	if err != nil {
		return nil, err
	}

	var res struct {
		Data Email `json:"data"`
		Email
	}
	if err := s.client.Perform(req, &res); err != nil {
		if req2, err2 := s.client.NewRequest(ctx, "GET", "/emails/"+emailId, nil); err2 == nil {
			if err3 := s.client.Perform(req2, &res); err3 == nil {
				err = nil
			}
		}
		if err != nil {
			return nil, err
		}
	}
	if res.Data.Id != "" {
		return &res.Data, nil
	}
	return &res.Email, nil
}
