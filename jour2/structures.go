package main

import "fmt"

type Config struct {
	ApiUrl    string
	Users     int
	Duration  int
	isRunning bool
}

func main() {
	config := Config{
		ApiUrl:    "http://loaclhost:8080",
		Users:     25,
		Duration:  60,
		isRunning: false,
	}

	fmt.Println(config.ApiUrl)
	fmt.Println(config.Users)
	fmt.Println(config.Duration)
	fmt.Println("Test en cours : ", config.isRunning)

	config.Users = 100

	fmt.Println("Nombre d'utilisateur mis à jour : ", config.Users)

	config.isRunning = true

	fmt.Print("Test en cours après modification : ", config.isRunning)

}
