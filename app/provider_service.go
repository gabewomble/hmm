package main

import (
	"app/db"
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"time"
)

type ProviderService struct{}

type ProviderInput struct {
	Name         string  `json:"name"`
	ProviderType string  `json:"providerType"`
	BaseURL      *string `json:"baseUrl"`
	APIKey       *string `json:"apiKey"`
}

type ProviderResponse struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	ProviderType string  `json:"providerType"`
	BaseURL      *string `json:"baseUrl"`
	APIKey       *string `json:"apiKey"`
	CreatedAt    int64   `json:"createdAt"`
	UpdatedAt    int64   `json:"updatedAt"`
}

func (s *ProviderService) CreateProvider(input ProviderInput) (ProviderResponse, error) {
	queries := db.New(db.DB())
	ctx := context.Background()

	now := time.Now().Unix()
	id := generateID()

	baseURL := sql.NullString{}
	if input.BaseURL != nil {
		baseURL = sql.NullString{String: *input.BaseURL, Valid: true}
	}

	apiKey := sql.NullString{}
	if input.APIKey != nil {
		apiKey = sql.NullString{String: *input.APIKey, Valid: true}
	}

	result, err := queries.CreateProvider(ctx, db.CreateProviderParams{
		ID:           id,
		Name:         input.Name,
		ProviderType: input.ProviderType,
		BaseUrl:      baseURL,
		ApiKey:       apiKey,
		CreatedAt:    now,
		UpdatedAt:    now,
	})
	if err != nil {
		return ProviderResponse{}, err
	}
	return toProviderResponse(result), nil
}

func (s *ProviderService) ListProviders() ([]ProviderResponse, error) {
	queries := db.New(db.DB())
	results, err := queries.ListProviders(context.Background())
	if err != nil {
		return nil, err
	}
	responses := make([]ProviderResponse, len(results))
	for i, r := range results {
		responses[i] = toProviderResponse(r)
	}
	return responses, nil
}

func (s *ProviderService) GetProvider(id string) (ProviderResponse, error) {
	queries := db.New(db.DB())
	result, err := queries.GetProvider(context.Background(), id)
	if err != nil {
		return ProviderResponse{}, err
	}
	return toProviderResponse(result), nil
}

func (s *ProviderService) DeleteProvider(id string) error {
	queries := db.New(db.DB())
	return queries.DeleteProvider(context.Background(), id)
}

func toProviderResponse(p db.Provider) ProviderResponse {
	r := ProviderResponse{
		ID:           p.ID,
		Name:         p.Name,
		ProviderType: p.ProviderType,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
	}
	if p.BaseUrl.Valid {
		r.BaseURL = &p.BaseUrl.String
	}
	if p.ApiKey.Valid {
		r.APIKey = &p.ApiKey.String
	}
	return r
}

func generateID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}
