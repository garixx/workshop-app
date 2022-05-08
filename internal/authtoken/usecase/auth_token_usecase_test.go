package usecase

import (
	"github.com/garixx/workshop-app/internal/domain"
	"testing"
	"time"
)

func TestAuthTokenUsecase_IsExpired(t *testing.T) {
	tests := []struct {
		name  string
		token domain.AuthToken
		wait  time.Duration
		want  bool
	}{
		{name: "current time is before expiredIn time", token: domain.AuthToken{CreatedAt: time.Now(), ExpiredIn: 10}, wait: 1 * time.Second, want: false},
		{name: "current time is after expiredIn time", token: domain.AuthToken{CreatedAt: time.Now(), ExpiredIn: 1}, wait: 2 * time.Second, want: true},
		{name: "current time is same as expiredIn time", token: domain.AuthToken{CreatedAt: time.Now(), ExpiredIn: 1}, wait: 1 * time.Second, want: true},
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
