package main

import (
	"net/http"
	"testing"

	"github.com/leonardoaraujodf/social/internal/store"
)

func TestGetUser(t *testing.T) {
	cfg := config{
		redisCfg: redisConfig{
			enabled: true,
		},
		chiLogger: false,
	}
	app := newTestApplication(t, cfg)
	mux := app.mount()
	testToken, err := app.authenticator.GenerateToken(nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Run("should not allow unauthenticated requests", func(t *testing.T) {
		// check for 401 code
		req, err := http.NewRequest(http.MethodGet, "/v1/users/1", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := executeRequest(req, mux)
		checkResponseCode(t, http.StatusUnauthorized, rr.Code)
	})
	t.Run("should allow authenticated requests", func(t *testing.T) {
		// check for 401 code
		req, err := http.NewRequest(http.MethodGet, "/v1/users/1", nil)
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Authorization", "Bearer "+testToken)
		rr := executeRequest(req, mux)
		checkResponseCode(t, http.StatusOK, rr.Code)
	})
}

func TestFollowUser(t *testing.T) {
	cfg := config{
		redisCfg: redisConfig{
			enabled: true,
		},
		chiLogger: false,
	}
	app := newTestApplication(t, cfg)
	mux := app.mount()
	testToken, err := app.authenticator.GenerateToken(nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("should not allow unauthenticated requests", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPut, "/v1/users/2/follow", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := executeRequest(req, mux)
		checkResponseCode(t, http.StatusUnauthorized, rr.Code)
	})

	t.Run("should follow user successfully", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPut, "/v1/users/2/follow", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Authorization", "Bearer "+testToken)
		rr := executeRequest(req, mux)
		checkResponseCode(t, http.StatusNoContent, rr.Code)
	})

	t.Run("should return 409 when already following", func(t *testing.T) {
		app.store.Followers = &store.MockFollowersStore{
			FollowErr: store.ErrConflict,
		}
		req, err := http.NewRequest(http.MethodPut, "/v1/users/2/follow", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Authorization", "Bearer "+testToken)
		rr := executeRequest(req, mux)
		checkResponseCode(t, http.StatusConflict, rr.Code)
	})
}
