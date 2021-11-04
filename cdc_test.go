package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

const port int32 = 4000

func startServer() {
	http.HandleFunc("/todos", GetToDos)
	http.HandleFunc("/todo", PostToDo)
	log.Fatalf("%v", http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
func TestVerifyContracts(t *testing.T) {
	godotenv.Load()
	go startServer()
	pact := &dsl.Pact{
		Provider:                 "Provider",
		DisableToolValidityCheck: true,
	}
	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL:            fmt.Sprintf("http://localhost:%v", port),
		BrokerURL:                  "https://rmustafayev.pactflow.io",
		PublishVerificationResults: true,
		ProviderVersion:            "1.0.0",
		BrokerToken:                os.Getenv("PACT_BROKER_TOKEN"),
	})
	if err != nil || t.Failed() {
		log.Fatalf("%v", err)
	}
}
