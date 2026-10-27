package main

import "fmt"

func verificationConfig(users int) {
	if users > 0 {
		fmt.Println("Configuration valide")
	} else {
		fmt.Println("Configuration invalide")
	}
}

func verifierCharge(users int) {
	if users <= 0 {
		fmt.Println("Configuartion invalide")
	} else if users >= 1 && users <= 20 {
		fmt.Println("Charge faible")
	} else {
		fmt.Println("Charge elevée")
	}
}

func main() {
	verificationConfig(0)
	verificationConfig(15)
	verifierCharge(0)
	verifierCharge(10)
	verifierCharge(50)
}
