package pinecone

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
)

type endpointI interface {
	buildURL(endpoint string) (*url.URL, error)
	newRequest(method string, u *url.URL, body interface{}) (*http.Request, error)
	doRequest(req *http.Request, v any) error
}

type endpoint struct {
	*Client
	EndpointPath string
}

func newEndpoint(c *Client, endpointPath string) *endpoint {
	e := &endpoint{
		Client:       c,
		EndpointPath: endpointPath,
	}
	return e
}

func (e *endpoint) buildURL(endpointPath string) (*url.URL, error) {
	u, err := url.Parse(endpointPath)
	if err != nil {
		return nil, err
	}
	sBaseUrl := e.BaseUrl
	if e.Environment != "" {
		sBaseUrl = fmt.Sprintf(apiTemplateUrl, e.Environment)
	}
	baseUrl, err := url.Parse(sBaseUrl)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(e.EndpointPath, u.Path)
	u.Path = path.Join(baseUrl.Path, u.Path)
	return baseUrl.ResolveReference(u), err
}

func (e *endpoint) doRequest(req *http.Request, v any) error {
	return e.Client.doRequest(req, v)
}

func (e *endpoint) newRequest(method string, u *url.URL, body interface{}) (*http.Request, error) {
	return e.Client.newRequest(method, u, body)
}
