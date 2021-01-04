package shortener

type RedirectRepository interface {
	FindByCode(code string) (*Redirect, error)
	FindByID(id string) (*Redirect, error)
	Save(redirect *Redirect) (*Redirect, error)
	Update(redirect *Redirect) (*Redirect, error)
}
