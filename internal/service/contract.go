package service

import (
	"context"

	"github.com/AugustSerenity/service_auth/internal/model"
)

type Storage interface {
	SaveToken(context.Context, *model.RefreshToken) error
	FindRefreshToken(context.Context, string, string) (*model.RefreshToken, error)
	MarkTokenUsed(context.Context, uint) error
}
