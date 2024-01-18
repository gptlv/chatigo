package server

import "github.com/gptlv/chatigo/server/internal/domain"

func (sr *serverRepository) UnregisterUser(u *domain.User) (*domain.User, error) {
	if _, ok := sr.users[u]; ok {
		delete(sr.users, u)
	}
}
