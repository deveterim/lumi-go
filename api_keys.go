package lumi

import "context"

type ApiKey struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Token     string `json:"token,omitempty"`
	Prefix    string `json:"prefix"`
	CreatedAt string `json:"created_at"`
}

type CreateApiKeyRequest struct {
	Name       string `json:"name"`
	Permission string `json:"permission,omitempty"`
}

type ApiKeysSvc struct {
	client *Client
}

func (s *ApiKeysSvc) Create(params *CreateApiKeyRequest) (*ApiKey, error) {
	return s.CreateWithContext(context.Background(), params)
}

func (s *ApiKeysSvc) CreateWithContext(ctx context.Context, params *CreateApiKeyRequest) (*ApiKey, error) {
	req, err := s.client.NewRequest(ctx, "POST", "/api-keys", params)
	if err != nil {
		return nil, err
	}
	var res ApiKey
	return &res, s.client.Perform(req, &res)
}

func (s *ApiKeysSvc) List() ([]ApiKey, error) {
	return s.ListWithContext(context.Background())
}

func (s *ApiKeysSvc) ListWithContext(ctx context.Context) ([]ApiKey, error) {
	req, err := s.client.NewRequest(ctx, "GET", "/api-keys", nil)
	if err != nil {
		return nil, err
	}
	var res struct {
		Data []ApiKey `json:"data"`
	}
	if err := s.client.Perform(req, &res); err != nil {
		return nil, err
	}
	return res.Data, nil
}

func (s *ApiKeysSvc) Remove(id string) error {
	return s.RemoveWithContext(context.Background(), id)
}

func (s *ApiKeysSvc) RemoveWithContext(ctx context.Context, id string) error {
	req, err := s.client.NewRequest(ctx, "DELETE", "/api-keys/"+id, nil)
	if err != nil {
		return err
	}
	return s.client.Perform(req, nil)
}
