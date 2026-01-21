package e2e

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"
)


func TestE2E_Car_Create(t *testing.T) {
	baseURL, cleanup := setupTestServer(t)
	defer cleanup()

	httpClient := &http.Client{}

	body := `{
			"brand": "Toyota",
			"model": "Camry",
			"year": 2020,
			"rent": 3000,
			"photo": "toyota.jpg"
	 	}`

	resp, err := httpClient.Post(
		baseURL+"/v1/cars/",
		"application/json",
		strings.NewReader(body),
	)
    if err != nil {
        t.Fatalf("failed to create request: %v", err)
    }

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected 201, got %d", resp.StatusCode)
	}

	var created struct {
		ID int64 `json:"id"`
	}

	json.NewDecoder(resp.Body).Decode(&created)

	if created.ID == 0 {
		t.Error("expected non-zero car ID")
	}

}