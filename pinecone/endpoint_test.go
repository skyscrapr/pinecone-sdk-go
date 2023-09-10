package pinecone

import (
	"testing"
)

func TestNewEndpoint(t *testing.T) {
	testEndpointPath := "testEndpointPath"
	testClient := NewClient("testapikey", "testenvironment")
	e := newEndpoint(testClient, testEndpointPath)
	if e.BaseUrl != apiTemplateUrl {
		t.Errorf("Endpoint baseUrl mismatch. Got %s. Want %s", e.BaseUrl, apiTemplateUrl)
	}
	if e.EndpointPath != testEndpointPath {
		t.Errorf("VendorsEndpoint EndpointPath mismatch. Got %s. Expected %s", e.EndpointPath, testEndpointPath)
	}
}
