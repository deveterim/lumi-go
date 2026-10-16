package mailifica

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

func TestWebhooksVerifySignature(t *testing.T) {
	secret := "whsec_test_secret"
	payload := []byte(`{"id":"evt_123","type":"email.delivered"}`)

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(payload)
	validSig := hex.EncodeToString(mac.Sum(nil))

	svc := &WebhooksSvc{}

	if !svc.VerifySignature(payload, validSig, secret) {
		t.Error("expected valid signature to pass")
	}

	if svc.VerifySignature(payload, "invalid_signature", secret) {
		t.Error("expected invalid signature to fail")
	}

	event, err := svc.ConstructEvent(payload, validSig, secret)
	if err != nil {
		t.Fatalf("unexpected error constructing event: %v", err)
	}

	if event.Id != "evt_123" {
		t.Errorf("expected evt_123, got %s", event.Id)
	}
}
