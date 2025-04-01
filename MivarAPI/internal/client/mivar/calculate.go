package mivar

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"mivar_robot_api/internal/client/dto"
)

func (c *Client) CalculatePath(ctx context.Context, in dto.CalculateRrequest) (*dto.CalculateResponse, error) {
	bodyBytes, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/ModelCalc", c.baseURL),
		bytes.NewReader(bodyBytes),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Add("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code %d: %s with request: %s", resp.StatusCode, string(body), string(bodyBytes))
	}

	var data dto.CalculateResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &data, nil
}
