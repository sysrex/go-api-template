package main

import (
	"fmt"

	"github.com/sysrex/go-api-template/internal/db"
)

// This will be responsable for the instantiaion and startup of the go app
func Run() error {
	fmt.Println("Start up with the error function")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}
	fmt.Println("connected and pinged the database")

	return nil
}

func main() {
	fmt.Println("Go Rest Template")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
