package main

import (
	"log"

	"github.com/krlspj/ex-bookings-hex-go/cmd/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
