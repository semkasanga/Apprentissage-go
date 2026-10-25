package main

import "fmt"

func main() {
	var users int = 10
	fmt.Println("Nombre d'utilisateurs : ", users)
	users = 25
	fmt.Println("Nombre d'utilisateurs après modification : ", users)

	var apiUrl string = "https://localhost:8080"
	var duration int = 30
	var isRunning bool = false

	fmt.Println("Url de l'API : ", apiUrl)
	fmt.Println("Durée du test : ", duration)
	fmt.Print("Test en cours : ", isRunning)
}
