package storage

import (
	"context"
	"lambda/model"
)

type Storage interface {
	SavePersona(ctx context.Context, persona *model.Persona) error
	GetPersonaByID(ctx context.Context, id string) (*model.Persona, error)
	GetAllPersonas(ctx context.Context) ([]*model.Persona, error)
}
