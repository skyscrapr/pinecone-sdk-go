package pinecone

import ()

const DatabasesEndpointPath = "/databases/"

type DatabasesEndpoint struct {
	*endpoint
}

type IndexMetric string

const (
	IndexMetricEuclidean  IndexMetric = "euclidean"
	IndexMetricCosine     IndexMetric = "cosine"
	IndexMetricDotProduct IndexMetric = "dotproduct"
)

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
	Dimension int    `json:"dimension"`
	Metric    string `json:"metric"`
	Pods      int    `json:"pods"`
	Replicas  int    `json:"replicas"`
	PodType   string `json:"pod_type"`
	Shards    int    `json:"shards"`
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

// name              string  required The name of the index to be created. The maximum length is 45 characters.
// dimension         integer required The dimensions of the vectors to be inserted in the index
// metric            string           The distance metric to be used for similarity search. You can use 'euclidean', 'cosine', or 'dotproduct'.
// pods              integer          The number of pods for the index to use,including replicas.
// replicas          integer          The number of replicas. Replicas duplicate your index. They provide higher availability and throughput.
// pod_type          string           The type of pod to use. One of s1, p1, or p2 appended with . and one of x1, x2, x4, or x8.
// metadata_config   object | null    Configuration for the behavior of Pinecone's internal metadata index. By default, all metadata is indexed; when metadata_config is present, only specified metadata fields are indexed. To specify metadata fields to index, provide a JSON object of the following form: {"indexed": ["example_metadata_field"]}
// source_collection string           The name of the collection to create an index from
type CreateIndexParams struct {
	Name             string            `json:"name"`
	Dimension        int               `json:"dimension"`
	Metric           IndexMetric       `json:"metric,omitempty"`
	Pods             *int              `json:"pods,omitempty"`
	Replicas         *int              `json:"replicas,omitempty"`
	PodType          *string           `json:"pod_type,omitempty"`
	MetadataConfig   map[string]string `json:"metadata_config,omitempty"`
	SourceCollection *string           `json:"source_collection,omitempty"`
}

// CreateIndex creates a Pinecone index.
//   - You can use it to specify the measure of similarity, the dimension of vectors to be stored in the index, the numbers of replicas to use, and more.
//
// API Reference: https://docs.pinecone.io/reference/create_index
func (e *DatabasesEndpoint) CreateIndex(params *CreateIndexParams) error {
	err := e.do(e, "POST", "", params, nil, nil)
	return err
}

// DescribeIndex gets description of an existing index.
// API Reference: https://docs.pinecone.io/reference/describe_index
func (e *DatabasesEndpoint) DescribeIndex(name string) (*Index, error) {
	var index Index
	err := e.do(e, "GET", name, nil, nil, &index)
	return &index, err
}

// DeleteIndex deletes an existing index.
// API Reference: https://docs.pinecone.io/reference/delete_index
func (e *DatabasesEndpoint) DeleteIndex(name string) error {
	err := e.do(e, "DELETE", name, nil, nil, nil)
	return err
}
