package cache

import (
	"context"

	"github.com/leonardoaraujodf/social/internal/store"
)

func NewMockStore() Storage {
	return Storage{
		Users: &MockUserStore{},
	}
}

type MockUserStore struct{}

func (m *MockUserStore) Get(ctx context.Context, id int64) (*store.User, error) {
	return nil, nil
}

func (m *MockUserStore) Set(ctx context.Context, u *store.User) error {
	return nil
}
