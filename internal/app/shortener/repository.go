package shortener

import "github.com/ahmetcancicek/go-url-shortener/internal/app/model"

type RedirectRepository interface {
	FindByCode(code string) (*model.Redirect, error)
	Save(redirect *model.Redirect) (*model.Redirect, error)
}
