package mivar

import (
	"context"
	"fmt"
	"net/http"

	"github.com/AlekSi/pointer"
)

func (c *Client) DeleteModel(ctx context.Context, modelID string) error {
	//Delete old model
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodDelete,
		fmt.Sprintf("%s/Models?modelID=%s", c.baseURL, modelID),
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to create delete request: %w", err)
	}

	req.Header.Add("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil && resp == nil {
		return fmt.Errorf("request failed: %w, response: %v", err, pointer.Get(resp).Body)
	}
	defer resp.Body.Close()

	return nil
}
