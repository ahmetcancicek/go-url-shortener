package shortener

type redirectService struct {
	redirectRepository RedirectRepository
}

func NewRedirectService(redirectRepository RedirectRepository) RedirectService {
	return &redirectService{
		redirectRepository,
	}
}

func (r redirectService) FindByCode(code string) (*Redirect, error) {
	panic("implement me")
}

func (r redirectService) FindByID(id string) (*Redirect, error) {
	panic("implement me")
}

func (r redirectService) Save(redirect *Redirect) (*Redirect, error) {
	panic("implement me")
}

func (r redirectService) Update(redirect *Redirect) (*Redirect, error) {
	panic("implement me")
}
