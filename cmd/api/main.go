package main

import (
	"fmt"
	"net/http"

	"github.com/kskumgk63/containized-firestore/internal/controller/api"
	"github.com/kskumgk63/containized-firestore/pkg/env"
)

func main() {
	port, err := env.Port()
	if err != nil {
		panic(err)
	}
	fmt.Println("========================")
	fmt.Printf(" 🚀 RUNNING ON %d🚀\n", port)
	fmt.Println("========================")
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), api.New()); err != nil {
		panic(err)
	}
}
