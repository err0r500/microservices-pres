package storer

import "github.com/err0r500/cleanarch-presentation/code/uc"

type Storer struct {
	Login string
	Pass  string
}

func (s Storer) NewUser(uc.User) (id int, err error) {
	return 0, nil
}

func (s Storer) SaveUser(uc.User) error {
	return nil
}

func (s Storer) GetUserByID(id string) (uc.User, error) {
	return uc.User{}, nil
}

type CredentialsGetter interface {
	GetLogin() string
	GetPassword() string
}

func NewInMemoryStorer(cg CredentialsGetter) Storer {
	return Storer{
		Login: cg.GetLogin(),
		Pass:  cg.GetPassword(),
	}
}
