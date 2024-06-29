package event

import (
	"context"
	"lambda/model"
)

type Publisher interface {
	PublishPersona(ctx context.Context, persona model.Persona) error
}
