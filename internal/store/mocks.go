package store

import (
	"context"
	"database/sql"
	"time"
)

func NewMockStore() Storage {
	return Storage{
		Users:     &MockUserStore{},
		Followers: &MockFollowersStore{},
	}
}

type MockUserStore struct{}

func (m *MockUserStore) Create(ctx context.Context, tx *sql.Tx, u *User) error {
	return nil
}

func (m *MockUserStore) GetByID(ctx context.Context, id int64) (*User, error) {
	return &User{ID: id, Username: "testuser", Email: "test@test.com"}, nil
}

func (m *MockUserStore) GetByEmail(ctx context.Context, email string) (*User, error) {
	return nil, nil
}

func (m *MockUserStore) CreateAndInvite(ctx context.Context, u *User, token string, exp time.Duration) error {
	return nil
}

func (m *MockUserStore) Activate(ctx context.Context, token string) error {
	return nil
}

func (m *MockUserStore) Delete(ctx context.Context, id int64) error {
	return nil
}

func (m *MockUserStore) CheckPassword(u *User, password string) (bool, error) {
	return false, nil
}

type MockFollowersStore struct {
	FollowErr   error
	UnfollowErr error
}

func (m *MockFollowersStore) FollowUser(ctx context.Context, userID, followerID int64) error {
	return m.FollowErr
}

func (m *MockFollowersStore) UnfollowUser(ctx context.Context, userID, followerID int64) error {
	return m.UnfollowErr
}
