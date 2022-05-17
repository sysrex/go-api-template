package main

import "fmt"

// This will be responsable for the instantiaion and startup of the go app
func Run() error {
	fmt.Println("Start up with the error function")
	return nil
}

func main() {
	fmt.Println("Go Rest Template")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
