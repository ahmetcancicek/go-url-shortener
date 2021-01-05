package shortener

type RedirectRepository interface {
	FindByCode(code string) (*Redirect, error)
	FindByID(id uint) (*Redirect, error)
	Save(redirect *Redirect) (*Redirect, error)
	Update(redirect *Redirect) (*Redirect, error)
}
