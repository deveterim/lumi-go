package mailifica

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	DefaultBaseURL = "https://api.mailifica.com/v1"
	Version        = "1.0.2"
)

type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client

	Emails   *EmailsSvc
	Batch    *BatchSvc
	Domains  *DomainsSvc
	ApiKeys  *ApiKeysSvc
	Webhooks *WebhooksSvc
}

func NewClient(apiKey string) *Client {
	if apiKey == "" {
		apiKey = os.Getenv("MAILIFICA_API_KEY")
	}

	baseURL := os.Getenv("MAILIFICA_BASE_URL")
	if baseURL == "" {
		baseURL = DefaultBaseURL
	}

	c := &Client{
		apiKey:  apiKey,
		baseURL: strings.TrimRight(baseURL, "/"),
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}

	c.Emails = &EmailsSvc{client: c}
	c.Batch = &BatchSvc{client: c}
	c.Domains = &DomainsSvc{client: c}
	c.ApiKeys = &ApiKeysSvc{client: c}
	c.Webhooks = &WebhooksSvc{}

	return c
}

func (c *Client) SetHttpClient(httpClient *http.Client) {
	c.httpClient = httpClient
}

func (c *Client) SetBaseURL(url string) {
	c.baseURL = strings.TrimRight(url, "/")
}

func (c *Client) NewRequest(ctx context.Context, method, path string, body any) (*http.Request, error) {
	if c.apiKey == "" {
		return nil, errors.New("mailifica: missing API key")
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		if err := json.NewEncoder(buf).Encode(body); err != nil {
			return nil, err
		}
	}

	url := c.baseURL + "/" + strings.TrimLeft(path, "/")
	req, err := http.NewRequestWithContext(ctx, method, url, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "mailifica-go/"+Version)

	return req, nil
}

func (c *Client) Perform(req *http.Request, v any) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var errResp ErrorResponse
		_ = json.NewDecoder(resp.Body).Decode(&errResp)
		errResp.StatusCode = resp.StatusCode
		if errResp.Message == "" {
			errResp.Message = http.StatusText(resp.StatusCode)
		}
		return &errResp
	}

	if v != nil && resp.StatusCode != http.StatusNoContent {
		return json.NewDecoder(resp.Body).Decode(v)
	}
	return nil
}
