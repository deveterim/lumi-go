package lumi

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
)

type WebhookEvent struct {
	Id        string         `json:"id"`
	Type      string         `json:"type"`
	CreatedAt string         `json:"created_at"`
	Data      map[string]any `json:"data"`
}

type WebhooksSvc struct{}

func (w *WebhooksSvc) VerifySignature(payload []byte, signature, secret string) bool {
	if signature == "" || secret == "" {
		return false
	}

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(payload)
	expectedMAC := hex.EncodeToString(mac.Sum(nil))

	return hmac.Equal([]byte(signature), []byte(expectedMAC))
}

func (w *WebhooksSvc) ConstructEvent(payload []byte, signature, secret string) (*WebhookEvent, error) {
	if !w.VerifySignature(payload, signature, secret) {
		return nil, errors.New("mailifica: invalid webhook signature")
	}

	var event WebhookEvent
	if err := json.Unmarshal(payload, &event); err != nil {
		return nil, err
	}
	return &event, nil
}
