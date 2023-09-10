package pinecone_test

import (
	"github.com/skyscrapr/pinecone-sdk-go/pinecone"
	"net/url"
)

const (
	test_api_key = "this-is-my-secure-apikey-do-not-steal!!"
	test_environment = "this-is-my-cloud"
)

func GetTestAuthToken() string {
	return test_api_key
}

func NewTestClient(ts *TestServer) *openai.Client {
	client := pinecone.NewClient(test_api_key, test_environment)
	if ts != nil {
		client.BaseURL, _ = url.Parse(ts.HTTPServer.URL)
	}
	return client
}
