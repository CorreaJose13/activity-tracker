package gemini

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var (
	errNoCandidates = errors.New("no candidates found")
	errNoResponse   = errors.New("no response found")
	errNoGemAPIKey  = errors.New("missing API key in environment variables")
)

func getAPIKey() (string, error) {
	key := os.Getenv("GEM_API_KEY")
	if key == "" {
		return "", errNoGemAPIKey
	}

	return key, nil
}

func extractFirstCandidate(candidates []*genai.Candidate) (string, error) {
	for _, candidate := range candidates {
		content := candidate.Content
		if content == nil || len(content.Parts) == 0 {
			continue
		}

		return fmt.Sprintf("%v", content.Parts[0]), nil
	}

	return "", errNoCandidates
}

func QueryGemini(prompt string) (string, error) {
	ctx := context.Background()

	apiKey, err := getAPIKey()
	if err != nil {
		return "", fmt.Errorf("API key retrieval error: %w", err)
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return "", fmt.Errorf("failed to create Gemini client: %w", err)
	}

	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", fmt.Errorf("failed to generate content: %w", err)
	}

	if resp == nil {
		fmt.Println("no response received from Gemini")
		return "", errNoResponse
	}

	candidates := resp.Candidates

	return extractFirstCandidate(candidates)
}
