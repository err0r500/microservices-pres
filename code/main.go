package main

import (
	"log"
	"time"

	"github.com/err0r500/cleanarch-presentation/code/config"
	storer "github.com/err0r500/cleanarch-presentation/code/implementation/inmemory.storer"
	"github.com/err0r500/cleanarch-presentation/code/uc"
)

func main() {

	credentialsGetter := config.NewCredentialsGetterFormEnv("LOGIN", "PASSWD")
	storer := storer.NewInMemoryStorer(credentialsGetter)
	interactor := uc.Interactor{ReadWriter: storer}

	err := interactor.CreateUser("matth", time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	log.Print(err)
}
