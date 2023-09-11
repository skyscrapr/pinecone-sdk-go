package pinecone

const CollectionsEndpointPath = "/collections"

type CollectionsEndpoint struct {
	*endpoint
}

// Collections Endpoint
func (c *Client) Collections() *CollectionsEndpoint {
	return &CollectionsEndpoint{newEndpoint(c, CollectionsEndpointPath)}
}

type Collection struct {
	Name   string `json:"name"`
	Size   int    `json:"size"`
	Status string `json:"status"`
}

// ListCollections returns a list of your Pinecone collections.
// API Reference: https://docs.pinecone.io/reference/list_collections
func (e *CollectionsEndpoint) ListCollections() ([]string, error) {
	var collections []string
	err := e.do(e, "GET", "", nil, nil, &collections)
	return collections, err
}

type CreateCollectionParams struct {
	Name   string `json:"name"`
	Source string `json:"source"`
}

// CreateCollection creates a Pinecone collection.
//   - Not supported by projects on the gcp-starter environment.
//
// API Reference: https://docs.pinecone.io/reference/create_collection
func (e *CollectionsEndpoint) CreateCollection(params *CreateCollectionParams) error {
	err := e.do(e, "POST", "", params, nil, nil)
	return err
}

// DescribeCollection gets description of an existing collection.
// API Reference: https://docs.pinecone.io/reference/describe_collection
func (e *CollectionsEndpoint) DescribeCollection(name string) (*Collection, error) {
	var collection Collection
	err := e.do(e, "GET", name, nil, nil, &collection)
	return &collection, err
}

// DeleteCollection deletes an existing collection.
// API Reference: https://docs.pinecone.io/reference/delete_collection
func (e *CollectionsEndpoint) DeleteCollection(name string) error {
	err := e.do(e, "DELETE", name, nil, nil, nil)
	return err
}
