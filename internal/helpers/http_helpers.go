package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type UserAttribute struct {
	Category     string    `json:"category"`
	Key          string    `json:"key"`
	KeywordValue []string  `json:"keyword_value,omitempty"`
	FloatValue   []float64 `json:"float_value,omitempty"`
	SourcedFrom  []string  `json:"sourced_from"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func GetUserAttributesFromEmail(email string) []UserAttribute {
	// Define the API URL
	url := fmt.Sprintf("http://hub.test:3000/api/headless/profile/attributes?email=%s", email)

	result, _ := FetchUserAttributes(url)

	// Print the response
	return result
}

func FetchUserAttributes(url string) ([]UserAttribute, error) {
	// Construct the URL with the email parameter

	// Make the HTTP request
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	// Check if the response status code is OK
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// Parse the JSON response
	var attributes []UserAttribute
	if err := json.Unmarshal(body, &attributes); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return attributes, nil
}
