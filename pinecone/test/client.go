package pinecone_test

import (
	"github.com/skyscrapr/pinecone-sdk-go/pinecone"
)

const (
	test_api_key     = "this-is-my-secure-apikey-do-not-steal!!"
	test_environment = "this-is-my-cloud"
)

func GetTestAuthToken() string {
	return test_api_key
}

func NewTestClient(ts *TestServer) *pinecone.Client {
	client := pinecone.NewClient(test_api_key, test_environment)
	if ts != nil {
		client.Environment = ""
		client.BaseUrl = ts.HTTPServer.URL
	}
	return client
}
