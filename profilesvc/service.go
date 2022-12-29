package profilesvc

import (
	"context"
	"errors"
	"sync"
)

var (
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
)

type Service interface {
	PostProfile(ctx context.Context, p Profile) error
}

type Profile struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Addresses []Address `json:"addresses, omitempty"`
}

type Address struct {
	ID       string `json:"id"`
	Location string `json:"location,omitempty"`
}

type inmemService struct {
	mtx sync.RWMutex
	m   map[string]Profile
}

func NewInmemService() Service {
	return &inmemService{
		m: map[string]Profile{},
	}
}

func (s *inmemService) PostProfile(ctx context.Context, p Profile) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if _, ok := s.m[p.ID]; ok {
		return ErrAlreadyExists
	}
	s.m[p.ID] = p
	return nil
}
