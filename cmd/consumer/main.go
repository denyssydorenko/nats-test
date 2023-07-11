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

	_, err = nc.Subscribe("gig", func(msg *nats.Msg) {
		// Manejar el mensaje recibido
		log.Printf("Mensaje recibido: %s", string(msg.Data))
	})
	if err != nil {
		log.Fatal(err)
	}

	// Esperar una señal para salir
	waitForExit()
}

func waitForExit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
