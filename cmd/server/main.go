package main

import "fmt"

// This will be responsable for the instantiaion and startup of the go app
func Run() error {
	fmt.Println("Start up with the error function")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go Rest Template")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
