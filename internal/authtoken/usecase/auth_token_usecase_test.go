package usecase

import (
	"github.com/garixx/workshop-app/internal/models"
	"github.com/garixx/workshop-app/internal/models/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestAuthTokenUsecase_IsExpired(t *testing.T) {
	tests := []struct {
		name  string
		token models.AuthToken
		wait  time.Duration
		want  bool
	}{
		{name: "current time is before expiredIn time", token: models.AuthToken{CreatedAt: time.Now(), ExpiredIn: 10}, wait: 1 * time.Second, want: false},
		{name: "current time is after expiredIn time", token: models.AuthToken{CreatedAt: time.Now(), ExpiredIn: 1}, wait: 2 * time.Second, want: true},
		{name: "current time is same as expiredIn time", token: models.AuthToken{CreatedAt: time.Now(), ExpiredIn: 1}, wait: 1 * time.Second, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := AuthTokenUsecase{
				userRepo: nil,
			}
			time.Sleep(tt.wait)
			if got := a.IsExpired(tt.token); got != tt.want {
				t.Errorf("IsExpired() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFetch(t *testing.T) {
	mockTokenRepo := mocks.NewAuthTokenRepository(t)

	expected := models.AuthToken{
		Login:     "me",
		Token:     "mocktokene5w6",
		CreatedAt: time.Now(),
		ExpiredIn: 10,
	}

	mockTokenRepo.On("FetchToken", mock.AnythingOfType("string")).Return(expected, nil).Once()

	tokenCase := NewAuthTokenUsecase(mockTokenRepo)
	token, err := tokenCase.ValidateToken("mocktokene5w6")

	assert.Equal(t, expected, token)
	assert.NoError(t, err)
}

func TestStore(t *testing.T) {
	mockTokenRepo := mocks.NewAuthTokenRepository(t)

	payload := models.AuthTokenParams{
		User: models.User{
			Login: "aaa",
		},
		Token: models.AuthToken{
			Token: "xxx",
		},
		ExpireIn: 10,
	}

	expected := models.AuthToken{
		Login:     "aaa",
		Token:     "aaaxxx",
		CreatedAt: time.Now(),
		ExpiredIn: 10,
	}

	//mockTokenRepo.On("StoreToken", mock.AnythingOfType("AuthTokenParams")).Return(expected, nil).Once()
	mockTokenRepo.On("StoreToken", payload).Return(expected, nil).Once()

	tokenCase := NewAuthTokenUsecase(mockTokenRepo)
	token, err := tokenCase.StoreToken(payload)

	assert.Equal(t, expected, token)
	assert.NoError(t, err)
}
