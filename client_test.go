package mailifica

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEmailsSend(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer ma_test_123" {
			t.Errorf("expected Authorization Bearer ma_test_123, got %s", r.Header.Get("Authorization"))
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("expected Content-Type application/json, got %s", r.Header.Get("Content-Type"))
		}
		if r.URL.Path != "/emails" {
			t.Errorf("expected path /emails, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(SendEmailResponse{
			Id:     "email_123",
			Status: "queued",
		})
	}))
	defer server.Close()

	client := NewClient("ma_test_123")
	client.SetBaseURL(server.URL)

	resp, err := client.Emails.Send(&SendEmailRequest{
		From:    "onboarding@empresa.ao",
		To:      []string{"dest@empresa.com"},
		Subject: "Boas-vindas",
		Html:    "<p>Olá!</p>",
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.Id != "email_123" {
		t.Errorf("expected email_123, got %s", resp.Id)
	}
}

func TestBatchSend(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/emails/batch" {
			t.Errorf("expected path /emails/batch, got %s", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(BatchSendResponse{
			Data: []SendEmailResponse{
				{Id: "email_1", Status: "queued"},
				{Id: "email_2", Status: "queued"},
			},
		})
	}))
	defer server.Close()

	client := NewClient("ma_test_123")
	client.SetBaseURL(server.URL)

	resp, err := client.Batch.Send([]*SendEmailRequest{
		{From: "a@a.ao", To: []string{"b@b.ao"}, Subject: "1", Html: "1"},
		{From: "a@a.ao", To: []string{"c@c.ao"}, Subject: "2", Html: "2"},
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(resp.Data) != 2 {
		t.Errorf("expected 2 batch items, got %d", len(resp.Data))
	}
}
