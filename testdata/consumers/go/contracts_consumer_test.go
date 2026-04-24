package consumer_test

import (
	"net/http"
	"testing"

	chillv4 "github.com/chill-institute/chill-contracts/gen/go/chill/v4"
	chillv4connect "github.com/chill-institute/chill-contracts/gen/go/chill/v4/chillv4connect"
)

func TestGeneratedContractsCompileForConsumer(t *testing.T) {
	request := &chillv4.GetUserProfileRequest{}
	if request == nil {
		t.Fatal("expected generated request type")
	}

	client := chillv4connect.NewUserServiceClient(http.DefaultClient, "https://example.com")
	if client == nil {
		t.Fatal("expected generated connect client")
	}
}
