package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"qinglong-envs/pkg/kv"
	"time"
)

type TokenService struct {
	store kv.Store
}

func NewTokenService(store kv.Store) *TokenService {
	return &TokenService{store: store}
}

type TokenPayload struct {
	UserId int64 `json:"userId"`
}

func (s *TokenService) CreateToken(ctx context.Context, userId int64) (string, error) {
	token := uuid.NewString()
	tp, err := json.Marshal(TokenPayload{UserId: userId})
	if err != nil {
		return "", err
	}
	err = s.store.SetWithTTL(
		fmt.Sprintf("token:%s", token),
		tp,
		24*time.Hour,
	)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *TokenService) GetTokenPayload(ctx context.Context, token string) (*TokenPayload, error) {
	tp, err := s.store.Get(fmt.Sprintf("token:%s", token))
	if err != nil {
		return nil, err
	}
	var payload TokenPayload
	err = json.Unmarshal(tp, &payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}
