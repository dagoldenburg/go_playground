package main

import (
	"GoReal/eventbus"
	"GoReal/subscribers"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	eb := eventbus.New()
	http.HandleFunc("/addEvent", func(writer http.ResponseWriter, request *http.Request) {
		addEvent(writer, request, eb)
	})

	http.HandleFunc("/selfCall", selfCall)

	go func() {
		printingSubscriber := &subscribers.PrintingSubscriber{Channel: make(chan interface{})}
		eb.Subscribe("addEvent", printingSubscriber)
	}()
	go func() {
		selfCallingSubscriber := &subscribers.SelfCallingSubscriber{Channel: make(chan interface{})}
		eb.Subscribe("addEvent", selfCallingSubscriber)
	}()

	startServer()
}

func selfCall(_ http.ResponseWriter, _ *http.Request) {
	fmt.Println("I called myself!!!... Why did I call myself?")
}

func startServer() {
	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func addEvent(response http.ResponseWriter, request *http.Request, eb *eventbus.EventBus) {
	if request.Method != "POST" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	byteData, err := io.ReadAll(request.Body)

	if err != nil {
		log.Println("Could not read body ", err)
		response.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	go func() {
		err := eb.Publish("addEvent", string(byteData))
		if err != nil {
			log.Println("Could not publish body ", err)
		}
	}()

	response.WriteHeader(http.StatusOK)
}
