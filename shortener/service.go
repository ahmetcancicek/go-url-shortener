package shortener

type RedirectService interface {
	FindByCode(code string) (*Redirect, error)
	Save(redirect *Redirect) (*Redirect, error)
}
