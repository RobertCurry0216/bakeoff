package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func BaseUrl() string {
	baseURL := os.Getenv("HUB_API_URL")
	if baseURL == "" {
		baseURL = "http://app:3000" // Default fallback
	}

	return baseURL
}

type UserAttributeRaw struct {
	Category     string    `json:"category"`
	Key          string    `json:"key"`
	KeywordValue []string  `json:"keyword_value,omitempty"`
	FloatValue   []float64 `json:"float_value,omitempty"`
}

type AttributeMetadataRaw struct {
	Name    string `json:"name"`
	Units   string `json:"units,omitempty"`
	Comment string `json:"comment"`
}

func GetUserAttributesFromEmail(email string) ([]AttributeData, error) {
	data, dataErr := FetchUserAttributes(email)
	metadata, metaErr := FetchAttributeMetadata(email)
	var err error
	if dataErr != nil || metaErr != nil {
		err = fmt.Errorf("meta: %s, data: %s", metaErr.Error(), dataErr.Error())
	}
	return MakeAttributes(data, metadata), err
}

func FetchUserAttributes(email string) ([]UserAttributeRaw, error) {
	// Construct the URL with the email parameter
	url := fmt.Sprintf("%s/api/headless/profile/attributes?email=%s", BaseUrl(), email)

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
	var attributes []UserAttributeRaw
	if err := json.Unmarshal(body, &attributes); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return attributes, nil
}

func FetchAttributeMetadata(email string) (map[string]AttributeMetadataRaw, error) {
	// Construct the URL with the email parameter
	url := fmt.Sprintf("%s/api/headless/profile/attribute_metadata?email=%s", BaseUrl(), email)

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
	var metadata map[string]AttributeMetadataRaw
	if err := json.Unmarshal(body, &metadata); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return metadata, nil
}
