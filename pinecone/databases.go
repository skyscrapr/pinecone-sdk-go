package pinecone

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

func (im IndexMetric) String() string {
	switch im {
	case IndexMetricEuclidean:
		return "euclidean"
	case IndexMetricCosine:
		return "cosine"
	case IndexMetricDotProduct:
		return "dotproduct"
	default:
		return ""
	}
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
	Name      string      `json:"name"`
	Dimension int         `json:"dimension"`
	Metric    IndexMetric `json:"metric"`
	Pods      int         `json:"pods"`
	Replicas  int         `json:"replicas"`
	PodType   string      `json:"pod_type"`
	Shards    int         `json:"shards"`
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

type CreateIndexParams struct {
	Name             string            `json:"name"`
	Dimension        int               `json:"dimension"`
	Metric           IndexMetric       `json:"metric"`
	Pods             int               `json:"pods"`
	Replicas         int               `json:"replicas"`
	PodType          string            `json:"pod_type"`
	MetadataConfig   map[string]string `json:"metadata_config,omitempty"`
	SourceCollection *string           `json:"source_collection,omitempty"`
}

type ConfigureIndexParams struct {
	Name     string `json:"name"`
	Replicas int    `json:"replicas"`
	PodType  string `json:"pod_type"`
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

// ConfigureIndex specifies the pod type and number of replicas for an index.
//   - Not supported by projects on the gcp-starter environment.
//
// API Reference: https://docs.pinecone.io/reference/configure_index
func (e *DatabasesEndpoint) ConfigureIndex(params *ConfigureIndexParams) error {
	err := e.do(e, "PATCH", "", params, nil, nil)
	return err
}

// DeleteIndex deletes an existing index.
// API Reference: https://docs.pinecone.io/reference/delete_index
func (e *DatabasesEndpoint) DeleteIndex(name string) error {
	err := e.do(e, "DELETE", name, nil, nil, nil)
	return err
}
