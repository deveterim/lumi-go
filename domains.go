package mailifica

import "context"

type DomainRecord struct {
	Record   string `json:"record"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Value    string `json:"value"`
	Status   string `json:"status"`
	Priority int    `json:"priority,omitempty"`
}

type Domain struct {
	Id        string         `json:"id"`
	Name      string         `json:"name"`
	Status    string         `json:"status"`
	Region    string         `json:"region,omitempty"`
	Records   []DomainRecord `json:"records"`
	CreatedAt string         `json:"created_at"`
}

type DomainsSvc struct {
	client *Client
}

func (s *DomainsSvc) Create(name string) (*Domain, error) {
	return s.CreateWithContext(context.Background(), name)
}

func (s *DomainsSvc) CreateWithContext(ctx context.Context, name string) (*Domain, error) {
	req, err := s.client.NewRequest(ctx, "POST", "/domains", map[string]string{"name": name})
	if err != nil {
		return nil, err
	}
	var res Domain
	return &res, s.client.Perform(req, &res)
}

func (s *DomainsSvc) List() ([]Domain, error) {
	return s.ListWithContext(context.Background())
}

func (s *DomainsSvc) ListWithContext(ctx context.Context) ([]Domain, error) {
	req, err := s.client.NewRequest(ctx, "GET", "/domains", nil)
	if err != nil {
		return nil, err
	}
	var res struct {
		Data []Domain `json:"data"`
	}
	if err := s.client.Perform(req, &res); err != nil {
		return nil, err
	}
	return res.Data, nil
}

func (s *DomainsSvc) Get(id string) (*Domain, error) {
	return s.GetWithContext(context.Background(), id)
}

func (s *DomainsSvc) GetWithContext(ctx context.Context, id string) (*Domain, error) {
	req, err := s.client.NewRequest(ctx, "GET", "/domains/"+id, nil)
	if err != nil {
		return nil, err
	}
	var res Domain
	return &res, s.client.Perform(req, &res)
}

func (s *DomainsSvc) Verify(id string) (*Domain, error) {
	return s.VerifyWithContext(context.Background(), id)
}

func (s *DomainsSvc) VerifyWithContext(ctx context.Context, id string) (*Domain, error) {
	req, err := s.client.NewRequest(ctx, "POST", "/domains/"+id+"/verify", nil)
	if err != nil {
		return nil, err
	}
	var res Domain
	return &res, s.client.Perform(req, &res)
}

func (s *DomainsSvc) Remove(id string) error {
	return s.RemoveWithContext(context.Background(), id)
}

func (s *DomainsSvc) RemoveWithContext(ctx context.Context, id string) error {
	req, err := s.client.NewRequest(ctx, "DELETE", "/domains/"+id, nil)
	if err != nil {
		return err
	}
	return s.client.Perform(req, nil)
}
