package pinecone

const IndexesEndpointPath = "/indexes/"

type IndexesEndpoint struct {
	*endpoint
}

type Index struct {
	Name      string      `json:"name"`
	Dimension int         `json:"dimension"`
	Metric    IndexMetric `json:"metric"`
	Host      string      `json:"host"`
	Status    IndexStatus `json:"status"`
	Spec      IndexSpec   `json:"spec"`
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

type IndexStatus struct {
	State IndexState `json:"state"`
	Ready bool       `json:"ready"`
}

type IndexState string

const (
	IndexStateInitializing IndexState = "Initializing"
	IndexStateScalingUp    IndexState = "ScalingUp"
	IndexStateScalingDown  IndexState = "ScalingDown"
	IndexStateTerminating  IndexState = "Terminating"
	IndexStateReady        IndexState = "Ready"
)

func (is IndexState) String() string {
	switch is {
	case IndexStateInitializing:
		return "Initializing"
	case IndexStateScalingUp:
		return "ScalingUp"
	case IndexStateScalingDown:
		return "ScalingDown"
	case IndexStateTerminating:
		return "Terminating"
	case IndexStateReady:
		return "Ready"
	default:
		return ""
	}
}

type IndexSpec struct {
	Serverless *IndexServerlessSpec `json:"serverless,omitempty"`
	Pod        *IndexPodSpec        `json:"pod,omitempty"`
}

type IndexServerlessSpec struct {
	Cloud  string `json:"cloud"`
	Region string `json:"region"`
}

type IndexPodSpec struct {
	Environment      string              `json:"environment"`
	Replicas         int                 `json:"replicas"`
	Shards           int                 `json:"shards"`
	PodType          string              `json:"pod_type"`
	Pods             int                 `json:"pods"`
	MetadataConfig   IndexMetadataConfig `json:"metadata_config"`
	SourceCollection string              `json:"source_collection,omitempty"`
}

type IndexMetadataConfig struct {
	Indexed []string `json:"indexed"`
}

// Databases Endpoint
func (c *Client) Indexes() *IndexesEndpoint {
	return &IndexesEndpoint{newEndpoint(c, IndexesEndpointPath)}
}

// ListIndexes returns a list of your Pinecone indexes.
// API Reference: https://docs.pinecone.io/reference/list_indexes
func (e *IndexesEndpoint) ListIndexes() ([]Index, error) {
	var indexes []Index
	err := e.do(e, "GET", "", nil, nil, &indexes)
	return indexes, err
}

type CreateIndexParams struct {
	Name      string      `json:"name"`
	Dimension int         `json:"dimension"`
	Metric    IndexMetric `json:"metric"`
	Spec      IndexSpec   `json:"spec"`
}

type ConfigureIndexParams struct {
	Spec IndexSpec `json:"spec"`
}

// CreateIndex creates a Pinecone index.
//   - You can use it to specify the measure of similarity, the dimension of vectors to be stored in the index, the numbers of replicas to use, and more.
//
// API Reference: https://docs.pinecone.io/reference/create_index
func (e *IndexesEndpoint) CreateIndex(params *CreateIndexParams) error {
	err := e.do(e, "POST", "", params, nil, nil)
	return err
}

// DescribeIndex gets description of an existing index.
// API Reference: https://docs.pinecone.io/reference/describe_index
func (e *IndexesEndpoint) DescribeIndex(name string) (*Index, error) {
	var index Index
	err := e.do(e, "GET", name, nil, nil, &index)
	return &index, err
}

// ConfigureIndex specifies the pod type and number of replicas for an index.
//   - Not supported by projects on the gcp-starter environment.
//
// API Reference: https://docs.pinecone.io/reference/configure_index
func (e *IndexesEndpoint) ConfigureIndex(name string, params *ConfigureIndexParams) error {
	err := e.do(e, "PATCH", name, params, nil, nil)
	return err
}

// DeleteIndex deletes an existing index.
// API Reference: https://docs.pinecone.io/reference/delete_index
func (e *IndexesEndpoint) DeleteIndex(name string) error {
	err := e.do(e, "DELETE", name, nil, nil, nil)
	return err
}
