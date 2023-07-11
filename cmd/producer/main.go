package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nats-io/nats.go"
)

func main() {
	// Conectar al servidor NATS
	nc, err := nats.Connect("nats://nats:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	message := []byte("Mensaje a publicar")

	err = nc.Publish("gig", message)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Mensaje publicado correctamente")
	// Esperar una señal para salir
	waitForExit()
}

func waitForExit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
