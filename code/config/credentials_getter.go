package config

import "os"

type CredentialsGetterFormEnv struct {
	LoginVarName    string
	PasswordVarName string
}

func NewCredentialsGetterFormEnv(loginVar, passVar string) CredentialsGetterFormEnv {
	return CredentialsGetterFormEnv{
		LoginVarName:    loginVar,
		PasswordVarName: passVar,
	}
}

func (cg CredentialsGetterFormEnv) GetLogin() string {
	return os.Getenv(cg.LoginVarName)
}

func (cg CredentialsGetterFormEnv) GetPassword() string {
	return os.Getenv(cg.PasswordVarName)
}
