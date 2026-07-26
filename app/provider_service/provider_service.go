package provider_service

import (
	"app/storage"
	"context"
	"time"
)

type ProviderService struct {
	store storage.ProviderStore
}

type ProviderInput struct {
	Name         string  `json:"name"`
	ProviderType string  `json:"providerType"`
	BaseURL      *string `json:"baseUrl"`
	APIKey       *string `json:"apiKey"`
}

type ProviderResponse struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	ProviderType string    `json:"providerType"`
	BaseURL      *string   `json:"baseUrl"`
	APIKey       *string   `json:"apiKey"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func New(store storage.ProviderStore) *ProviderService {
	return &ProviderService{store: store}
}

func (s *ProviderService) CreateProvider(input ProviderInput) (ProviderResponse, error) {
	result, err := s.store.CreateProvider(context.Background(), storage.CreateProviderParams{
		Name:         input.Name,
		ProviderType: input.ProviderType,
		BaseURL:      input.BaseURL,
		APIKey:       input.APIKey,
	})
	if err != nil {
		return ProviderResponse{}, err
	}
	return toResponse(result), nil
}

func (s *ProviderService) GetProvider(id string) (ProviderResponse, error) {
	result, err := s.store.GetProvider(context.Background(), id)
	if err != nil {
		return ProviderResponse{}, err
	}
	return toResponse(result), nil
}

func (s *ProviderService) ListProviders() ([]ProviderResponse, error) {
	results, err := s.store.ListProviders(context.Background())
	if err != nil {
		return nil, err
	}
	responses := make([]ProviderResponse, len(results))
	for i, r := range results {
		responses[i] = toResponse(r)
	}
	return responses, nil
}

func (s *ProviderService) UpdateProvider(id string, input ProviderInput) (ProviderResponse, error) {
	result, err := s.store.UpdateProvider(context.Background(), storage.UpdateProviderParams{
		ID:           id,
		Name:         input.Name,
		ProviderType: input.ProviderType,
		BaseURL:      input.BaseURL,
		APIKey:       input.APIKey,
	})
	if err != nil {
		return ProviderResponse{}, err
	}
	return toResponse(result), nil
}

func (s *ProviderService) DeleteProvider(id string) error {
	return s.store.DeleteProvider(context.Background(), id)
}

func toResponse(p storage.Provider) ProviderResponse {
	return ProviderResponse{
		ID:           p.ID,
		Name:         p.Name,
		ProviderType: p.ProviderType,
		BaseURL:      p.BaseURL,
		APIKey:       p.APIKey,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
	}
}
