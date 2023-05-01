package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func writeToDotEnv() {
	env, _ := godotenv.Unmarshal("test=test1")
	err := godotenv.Write(env, ".env")
	if err != nil {
		log.Println("There was an error writing to the dotenv file")
	}
}

func getDotEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading environment variables file")
	}
	name := os.Getenv("test")
	fmt.Printf(name)
}
