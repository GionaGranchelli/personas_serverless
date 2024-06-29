package storage

import (
	"context"
	"lambda/model"
)

type Storage interface {
	SavePersona(ctx context.Context, persona *model.Persona) error
}
