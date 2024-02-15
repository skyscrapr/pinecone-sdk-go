package pinecone

import (
	"testing"
)

func TestBadApiKey(t *testing.T) {
	client := NewClient("bad-api-key", "gcp-starter")
	_, err := client.Indexes().ListIndexes()
	if err != nil && err.Error() != "403 Forbidden: Invalid API key" {
		t.Errorf("Expected Error '403 Forbidden: Invalid API key'. Got %s", err)
	}
}
