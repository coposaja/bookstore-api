package main

import (
	"fmt"
	"log"
	"os"

	"github.com/coposaja/bookstore-api/src/db/mysql"
	"github.com/coposaja/bookstore-api/src/startup"
	"github.com/joho/godotenv"
)

func init() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			log.Println("Please contact your administrator...")
			log.Println("Press enter to exit...")
			fmt.Scanln()
			os.Exit(1)
		}
	}()
	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file")
	}

	mysql.Connect()
}

func main() {
	startup.StartApp()
}
