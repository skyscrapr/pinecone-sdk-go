package openai_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/skyscrapr/pinecone-sdk-go/pinecone/test"
)

func TestListIndexes(t *testing.T) {
	ts := pinecone_test.NewTestServer()
	ts.RegisterHandler("/databases/", func(w http.ResponseWriter, _ *http.Request) {
		resBytes, _ := json.Marshal([]string{
			"index_1",
			"index_2",
		})
		fmt.Fprintln(w, string(resBytes))
	})
	ts.HTTPServer.Start()
	defer ts.HTTPServer.Close()

	client := pinecone_test.NewTestClient(ts)
	_, err := client.Databases().ListIndexes()
	t.Helper()
	if err != nil {
		t.Error(err, "CreateFineTuningJob error")
		t.Fail()
	}
}
