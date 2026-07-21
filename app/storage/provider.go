package storage

import "context"

type Provider struct {
	ID           string
	Name         string
	ProviderType string
	BaseURL      *string
	APIKey       *string
	CreatedAt    int64
	UpdatedAt    int64
}

type CreateProviderParams struct {
	Name         string
	ProviderType string
	BaseURL      *string
	APIKey       *string
}

type UpdateProviderParams struct {
	ID           string
	Name         string
	ProviderType string
	BaseURL      *string
	APIKey       *string
}

type ProviderStore interface {
	CreateProvider(ctx context.Context, params CreateProviderParams) (Provider, error)
	GetProvider(ctx context.Context, id string) (Provider, error)
	ListProviders(ctx context.Context) ([]Provider, error)
	UpdateProvider(ctx context.Context, params UpdateProviderParams) (Provider, error)
	DeleteProvider(ctx context.Context, id string) error
}
