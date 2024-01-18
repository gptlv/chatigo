package server

import "github.com/gptlv/chatigo/server/internal/domain"

func (sr *serverRepository) RegisterUser(u *domain.User) (*domain.User, error) {
	sr.users[u] = true
	return
}
