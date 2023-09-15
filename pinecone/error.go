package pinecone

import (
	"fmt"
	"io"
	"net/http"
)

func (c *Client) handleHTTPErrorResp(resp *http.Response) error {
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	reqErr := &HTTPError{
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Message:    string(data),
	}
	return reqErr
}

// HTTPError provides informations about generic HTTP errors.
type HTTPError struct {
	StatusCode int
	Status     string
	Message    string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("error, status code: %d, message: %s", e.StatusCode, e.Message)
}
