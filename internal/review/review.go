package review

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/vossenwout/crev/internal/files"
)

type ReviewInput struct {
	Code string `json:"code"`
}

type ReviewOutput struct {
	Review string `json:"review"`
}

const reviewURL = "https://reviewcode-qcgl4feadq-uc.a.run.app"

func prepareRequest(codeToReview string, apiKey string) (*http.Request, error) {
	input := ReviewInput{
		Code: codeToReview,
	}
	jsonData, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", reviewURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	// Set the request header to specify JSON format
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api-key", apiKey)

	return req, nil
}

func sendRequest(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending review request to %s: %v", reviewURL, err)
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("Error: received status code %d: %s", resp.StatusCode, string(body))
		if resp.StatusCode == http.StatusUnauthorized {
			log.Fatalf("Unauthorized: you have provided an invalid CREV API key.")
		} else {
			log.Fatalf("Failed to review code: status code %d", resp.StatusCode)
		}
		return nil, err
	}
	return resp, nil
}

func saveReviewToFile(output ReviewOutput) error {
	err := files.SaveStringToFile(output.Review, "crev-review.md")
	if err != nil {
		return err
	}
	log.Printf("Successfully saved code review to crev-review.md")
	return nil
}

func Review(codeToReview string, apiKey string) {
	log.Printf("Reviewing code please wait...")

	// Prepare the request to review the code
	req, err := prepareRequest(codeToReview, apiKey)
	if err != nil {
		log.Fatalf("Error preparing review request: %v", err)
	}

	// Send the request to review the code
	resp, err := sendRequest(req)
	if err != nil {
		log.Fatalf("Error sending review request: %v", err)
	}
	defer resp.Body.Close()

	// Decode the response
	var output ReviewOutput
	err = json.NewDecoder(resp.Body).Decode(&output)
	if err != nil {
		log.Fatalf("Error decoding review response: %v", err)
	}

	// Save the review to a file
	err = saveReviewToFile(output)
	if err != nil {
		log.Fatalf("Error saving review to file: %v", err)
	}

}
