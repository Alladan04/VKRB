package mivar

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"mivar_robot_api/internal/client/dto"
)

func (c *Client) AddModel(ctx context.Context, in dto.AddModelRequest) (*dto.AddModelResponse, error) {
	bodyBytes, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/Models", c.baseURL),
		bytes.NewReader(bodyBytes),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Add("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil && resp == nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	var data dto.AddModelResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &data, nil
}
