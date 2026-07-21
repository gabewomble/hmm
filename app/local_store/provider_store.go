package local_store

import (
	"app/storage"
	"context"
	"database/sql"
	"time"
)

func (s *LocalStore) CreateProvider(ctx context.Context, params storage.CreateProviderParams) (storage.Provider, error) {
	now := time.Now().Unix()
	result, err := s.queries.CreateProvider(ctx, CreateProviderParams{
		Name:         params.Name,
		ProviderType: params.ProviderType,
		BaseUrl:      toNullString(params.BaseURL),
		ApiKey:       toNullString(params.APIKey),
		CreatedAt:    now,
		UpdatedAt:    now,
	})
	if err != nil {
		return storage.Provider{}, err
	}
	return toStorageProvider(result), nil
}

func (s *LocalStore) GetProvider(ctx context.Context, id string) (storage.Provider, error) {
	result, err := s.queries.GetProvider(ctx, id)
	if err != nil {
		return storage.Provider{}, err
	}
	return toStorageProvider(result), nil
}

func (s *LocalStore) ListProviders(ctx context.Context) ([]storage.Provider, error) {
	results, err := s.queries.ListProviders(ctx)
	if err != nil {
		return nil, err
	}
	providers := make([]storage.Provider, len(results))
	for i, r := range results {
		providers[i] = toStorageProvider(r)
	}
	return providers, nil
}

func (s *LocalStore) UpdateProvider(ctx context.Context, params storage.UpdateProviderParams) (storage.Provider, error) {
	result, err := s.queries.UpdateProvider(ctx, UpdateProviderParams{
		ID:           params.ID,
		Name:         params.Name,
		ProviderType: params.ProviderType,
		BaseUrl:      toNullString(params.BaseURL),
		ApiKey:       toNullString(params.APIKey),
		UpdatedAt:    time.Now().Unix(),
	})
	if err != nil {
		return storage.Provider{}, err
	}
	return toStorageProvider(result), nil
}

func (s *LocalStore) DeleteProvider(ctx context.Context, id string) error {
	return s.queries.DeleteProvider(ctx, id)
}

func toNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{}
	}
	return sql.NullString{String: *s, Valid: true}
}

func toStorageProvider(p Provider) storage.Provider {
	return storage.Provider{
		ID:           p.ID,
		Name:         p.Name,
		ProviderType: p.ProviderType,
		BaseURL:      fromNullString(p.BaseUrl),
		APIKey:       fromNullString(p.ApiKey),
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
	}
}

func fromNullString(ns sql.NullString) *string {
	if !ns.Valid {
		return nil
	}
	return &ns.String
}
