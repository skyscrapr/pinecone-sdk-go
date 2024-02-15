package pinecone_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/skyscrapr/pinecone-sdk-go/pinecone"
	pinecone_test "github.com/skyscrapr/pinecone-sdk-go/pinecone/test"
)

func TestListIndexes(t *testing.T) {
	ts := pinecone_test.NewTestServer()
	ts.RegisterHandler("/indexes", func(w http.ResponseWriter, _ *http.Request) {
		resBytes, _ := json.Marshal([]string{
			"index_1",
			"index_2",
		})
		fmt.Fprintln(w, string(resBytes))
	})
	ts.HTTPServer.Start()
	defer ts.HTTPServer.Close()

	client := pinecone_test.NewTestClient(ts)
	_, err := client.Indexes().ListIndexes()
	t.Helper()
	if err != nil {
		t.Error(err, "ListIndexes error")
		t.Fail()
	}
}

func TestDescribeIndex(t *testing.T) {
	ts := pinecone_test.NewTestServer()
	ts.RegisterHandler("/indexes/test_index", func(w http.ResponseWriter, _ *http.Request) {
		resBytes, _ := json.Marshal(pinecone.Index{
			Name:      "test_index",
			Dimension: 512,
			Metric:    pinecone.IndexMetricCosine,
			Spec: pinecone.IndexSpec{
				Pod: &pinecone.IndexPodSpec{
					Environment: "us-west4-gcp",
					Replicas:    1,
					PodType:     "starter",
				},
			},
			Status: pinecone.IndexStatus{
				Ready: true,
				State: "Ready",
			},
		})
		fmt.Fprintln(w, string(resBytes))
	})
	ts.HTTPServer.Start()
	defer ts.HTTPServer.Close()

	client := pinecone_test.NewTestClient(ts)
	_, err := client.Indexes().DescribeIndex("test_index")
	t.Helper()
	if err != nil {
		t.Error(err, "DescribeIndex error")
		t.Fail()
	}
}
