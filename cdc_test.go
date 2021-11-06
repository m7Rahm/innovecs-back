package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"testing"

	"github.com/joho/godotenv"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

const port int32 = 4000

func startServer() {
	http.HandleFunc("/api/todos", GetToDos)
	http.HandleFunc("/api/todo", PostToDo)
	log.Fatalf("%v", http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
func TestVerifyContracts(t *testing.T) {
	godotenv.Load()
	go startServer()
	gitSha, _ := exec.Command("git", "rev-parse").Output()
	branch, _ := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()

	provider_tags := make([]string, 1)
	provider_tags = append(provider_tags, string(branch))
	pact := &dsl.Pact{
		Provider:                 "Provider",
		DisableToolValidityCheck: true,
	}
	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL:            fmt.Sprintf("http://localhost:%v", port),
		BrokerURL:                  "https://rmustafayev.pactflow.io",
		PublishVerificationResults: true,
		ProviderTags:               provider_tags,
		ProviderVersion:            fmt.Sprintf("0.0.%v", gitSha),
		BrokerToken:                os.Getenv("PACT_BROKER_TOKEN"),
	})
	if err != nil || t.Failed() {
		log.Fatalf("%v", err)
	}
}
