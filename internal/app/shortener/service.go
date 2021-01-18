package shortener

import (
	"context"
	"github.com/ahmetcancicek/go-url-shortener/internal/app/model"
)

type RedirectService interface {
	FindByCode(ctx context.Context, code string) (*model.Redirect, error)
	Save(ctx context.Context, redirect *model.Redirect) (*model.Redirect, error)
}
