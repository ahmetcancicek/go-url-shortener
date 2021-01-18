package service

import (
	"github.com/ahmetcancicek/go-url-shortener/internal/app/model"
	"github.com/ahmetcancicek/go-url-shortener/internal/app/shortener"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"time"
)

type redirectService struct {
	redirectRepository shortener.RedirectRepository
}

func NewRedirectService(redirectRepository shortener.RedirectRepository) shortener.RedirectService {
	return &redirectService{
		redirectRepository,
	}
}

func (r redirectService) FindByCode(code string) (*model.Redirect, error) {
	return r.redirectRepository.FindByCode(code)
}

func (r *redirectService) Save(redirect *model.Redirect) (*model.Redirect, error) {
	validate := validator.New()
	if err := validate.Struct(redirect); err != nil {
		return redirect, errors.Wrap(err, "service.Redirect.Save")
	}
	redirect.Click = 0
	redirect.CreatedAt = time.Now()
	return r.redirectRepository.Save(redirect)
}
