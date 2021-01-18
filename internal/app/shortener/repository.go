package shortener

type RedirectRepository interface {
	FindByCode(code string) (*Redirect, error)
	Save(redirect *Redirect) (*Redirect, error)
}
