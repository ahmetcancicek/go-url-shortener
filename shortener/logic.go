package shortener

import (
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"time"
)

type redirectService struct {
	redirectRepository RedirectRepository
}

func NewRedirectService(redirectRepository RedirectRepository) RedirectService {
	return &redirectService{
		redirectRepository,
	}
}

func (r redirectService) FindByCode(code string) (*Redirect, error) {
	return r.redirectRepository.FindByCode(code)
}

func (r *redirectService) Save(redirect *Redirect) (*Redirect, error) {
	validate := validator.New()
	if err := validate.Struct(redirect); err != nil {
		return redirect, errors.Wrap(err, "service.Redirect.Save")
	}
	redirect.Click = 0
	redirect.CreatedAt = time.Now()
	return r.redirectRepository.Save(redirect)
}
