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
	return fmt.Errorf(resp.Status + ": " + string(data))
}
