package pinecone

import ()

const DatabasesEndpointPath = "/databases/"

type DatabasesEndpoint struct {
	*endpoint
}

// Databases Endpoint
func (c *Client) Databases() *DatabasesEndpoint {
	return &DatabasesEndpoint{newEndpoint(c, DatabasesEndpointPath)}
}

type Index struct {
	Database Database `json:"database"`
	Status   Status   `json:"status"`
}

type Database struct {
	Name      string `json:"name"`
	Metric    string `json:"metric"`
	Dimension int    `json:"dimension"`
	Replicas  int    `json:"replicas"`
	Shards    int    `json:"shards"`
	Pods      int    `json:"pods"`
	PodType   string `json:"pod_type"`
}

type Status struct {
	Waiting []interface{} `json:"waiting"`
	Crashed []interface{} `json:"crashed"`
	Host    string        `json:"host"`
	Port    int           `json:"port"`
	State   string        `json:"state"`
	Ready   bool          `json:"ready"`
}

// ListIndexes returns a list of your Pinecone indexes.
// API Reference: https://docs.pinecone.io/reference/list_indexes
func (e *DatabasesEndpoint) ListIndexes() ([]string, error) {
	var indexes []string
	err := e.do(e, "GET", "", nil, nil, &indexes)
	return indexes, err
}
