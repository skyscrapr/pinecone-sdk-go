package pinecone_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/skyscrapr/pinecone-sdk-go/pinecone"
	"github.com/skyscrapr/pinecone-sdk-go/pinecone/test"
)

func TestListIndexes(t *testing.T) {
	ts := pinecone_test.NewTestServer()
	ts.RegisterHandler("/databases", func(w http.ResponseWriter, _ *http.Request) {
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
		t.Error(err, "ListIndexes error")
		t.Fail()
	}
}

func TestDescribeIndex(t *testing.T) {
	ts := pinecone_test.NewTestServer()
	ts.RegisterHandler("/databases/test_index", func(w http.ResponseWriter, _ *http.Request) {
		resBytes, _ := json.Marshal(pinecone.Index{
			Database: pinecone.Database{
				Name:      "test_index",
				Dimension: 512,
				Metric:    pinecone.IndexMetricCosine,
				Pods:      1,
				Replicas:  1,
				PodType:   "starter",
			},
			Status: pinecone.Status{
				Ready: true,
				State: "Ready",
			},
		})
		fmt.Fprintln(w, string(resBytes))
	})
	ts.HTTPServer.Start()
	defer ts.HTTPServer.Close()

	client := pinecone_test.NewTestClient(ts)
	_, err := client.Databases().DescribeIndex("test_index")
	t.Helper()
	if err != nil {
		t.Error(err, "DescribeIndex error")
		t.Fail()
	}
}
