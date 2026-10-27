package main

import "fmt"

func afficherBienvenue() {
	fmt.Println("Bienvenue dans API Load Tester")
}

func afficherObjectif() {
	fmt.Println("Notre objectif est de tester les performances d'une API")
}

func afficherUtilisateurs(users int) {
	fmt.Println("Nombre d'utilisateurs : ", users)
}

func calculerTotal(users int, requests int) int {
	return users * requests
}

func main() {
	afficherBienvenue()
	afficherObjectif()
	afficherUtilisateurs(10)
	afficherUtilisateurs(50)
	total := calculerTotal(10, 5)
	fmt.Println("Total :", total)
}
