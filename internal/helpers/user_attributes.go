package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// UserAttribute represents a single attribute from the API response
type UserAttribute struct {
	Category     string    `json:"category"`
	Key          string    `json:"key"`
	KeywordValue []string  `json:"keyword_value,omitempty"`
	FloatValue   []float64 `json:"float_value,omitempty"`
	SourcedFrom  []string  `json:"sourced_from"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// GetUserAttributesFromEmail fetches user attributes from the API using the provided email
func GetUserAttributesFromEmail(email string) []UserAttribute {
	// This is a mock implementation
	// In a real application, you would make an HTTP request to your API
	
	// Example data based on the provided JSON structure
	mockData := []UserAttribute{
		{
			Category:     "consumer",
			Key:          "au.com.demov2.email",
			KeywordValue: []string{email},
			SourcedFrom:  []string{"internet solution"},
			UpdatedAt:    time.Now(),
		},
		{
			Category:   "consumer",
			Key:        "au.com.demov2.transaction.avg_spend_per_order",
			FloatValue: []float64{90.06},
			SourcedFrom: []string{"static"},
			UpdatedAt:  time.Now().Add(-24 * time.Hour),
		},
		{
			Category:   "consumer",
			Key:        "au.com.demov2.transaction.total_spend",
			FloatValue: []float64{38.95},
			SourcedFrom: []string{"Managed"},
			UpdatedAt:  time.Now().Add(-48 * time.Hour),
		},
		{
			Category:   "consumer",
			Key:        "au.com.demov2.transaction.average_discount_level",
			FloatValue: []float64{97},
			SourcedFrom: []string{"encompassing"},
			UpdatedAt:  time.Now().Add(-72 * time.Hour),
		},
		{
			Category:   "consumer",
			Key:        "au.com.demov2.transaction.avg_spend_per_product",
			FloatValue: []float64{19.82},
			SourcedFrom: []string{"Stand-alone"},
			UpdatedAt:  time.Now().Add(-96 * time.Hour),
		},
	}

	return mockData
}

// FetchUserAttributes makes a real HTTP request to fetch user attributes
// This would be used in a production environment instead of the mock function
func FetchUserAttributes(email string, apiURL string) ([]UserAttribute, error) {
	// Construct the URL with the email parameter
	url := fmt.Sprintf("%s/user/%s", apiURL, email)
	
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
