package kv

import (
	"log"
	"os"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azlog "github.com/Azure/azure-sdk-for-go/sdk/azcore/log"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

var (
	credential azcore.TokenCredential
)

// https://pkg.go.dev/testing#hdr-Main

func TestMain(m *testing.M) {
	// Set log to output to the console
	azlog.SetListener(func(e azlog.Event, msg string) {
		log.Printf("%s: %s\n", e, msg) // printing log out to the console
	})

	// Includes only requests and responses in credential logs
	azlog.SetEvents(azlog.EventRequest, azlog.EventResponse, azlog.EventRetryPolicy, azlog.EventLRO)

	var err error
	credential, err = azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}

func TestListKeys(t *testing.T) {
	ListKeys(credential)
}

func TestListSecrets(t *testing.T) {
	ListSecrets(credential)
}

func TestGetSecrets(t *testing.T) {
	_, err := GetSecrets(credential, "server")
	if err != nil {
		log.Fatal(err)
	}
}

func TestGetSecrets2(t *testing.T) {
	_, err := GetSecrets(credential, "cert1")
	if err != nil {
		log.Fatal(err)
	}
}
