package domain

import "context"

// !!! 講解1 !!!
// Digimon ...
type Digimon struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

// !!! 講解2 !!!
// DigimonRepository ...
type DigimonRepository interface {
	GetByID(ctx context.Context, id string) (*Digimon, error)
	Store(ctx context.Context, d *Digimon) error
	UpdateStatus(ctx context.Context, d *Digimon) error
}

// !!! 講解3 !!!
// DigimonUsecase ..
type DigimonUsecase interface {
	GetByID(ctx context.Context, id string) (*Digimon, error)
	Store(ctx context.Context, d *Digimon) error
	UpdateStatus(ctx context.Context, d *Digimon) error
}
