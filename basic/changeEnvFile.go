package main

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main1() {
	err := godotenv.Load()
	if err != nil {
		log.Print(err)
	}
	dns := os.Getenv("dns")
	host := os.Getenv("host")
	r := strings.NewReplacer("<host>", host, "<port>", "5432", "<dbname>", "asd")
	out := r.Replace(dns)
	log.Print(out)
}
