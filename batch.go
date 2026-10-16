package lumi

import "context"

type BatchSvc struct {
	client *Client
}

type BatchSendResponse struct {
	Data []SendEmailResponse `json:"data"`
}

func (s *BatchSvc) Send(params []*SendEmailRequest) (*BatchSendResponse, error) {
	return s.SendWithContext(context.Background(), params)
}

func (s *BatchSvc) SendWithContext(ctx context.Context, params []*SendEmailRequest) (*BatchSendResponse, error) {
	items := make([]map[string]any, len(params))
	for i, p := range params {
		m := map[string]any{
			"from":    p.From,
			"subject": p.Subject,
		}
		if len(p.To) == 1 {
			m["to"] = p.To[0]
		} else if len(p.To) > 1 {
			m["to"] = p.To
		}
		if p.Html != "" {
			m["html"] = p.Html
		}
		if p.Text != "" {
			m["text"] = p.Text
		}
		if p.ReplyTo != "" {
			m["reply_to"] = p.ReplyTo
		}
		if len(p.Cc) > 0 {
			m["cc"] = p.Cc
		}
		if len(p.Bcc) > 0 {
			m["bcc"] = p.Bcc
		}
		if len(p.Headers) > 0 {
			m["headers"] = p.Headers
		}
		if len(p.Attachments) > 0 {
			m["attachments"] = p.Attachments
		}
		if len(p.Tags) > 0 {
			m["tags"] = p.Tags
		}
		if p.ScheduledAt != "" {
			m["scheduled_at"] = p.ScheduledAt
		}
		items[i] = m
	}
	payload := map[string]any{
		"emails": items,
	}
	req, err := s.client.NewRequest(ctx, "POST", "/emails/batch", payload)
	if err != nil {
		return nil, err
	}

	var res BatchSendResponse
	if err := s.client.Perform(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
