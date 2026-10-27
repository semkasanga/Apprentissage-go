package main

import (
	"errors"
	"fmt"
)

func verifierDuree(duration int) error {
	if duration <= 0 {
		return errors.New("La durée doit etre positive")
	}

	return nil
}

func main() {
	err := verifierDuree(0)
	if err != nil {
		fmt.Println("Erreur :", err)
	} else {
		fmt.Println("Durée valide")
	}

	err = verifierDuree(30)
	if err != nil {
		fmt.Println("Erreur :", err)
	} else {
		fmt.Println("Durée valide")
	}
}
