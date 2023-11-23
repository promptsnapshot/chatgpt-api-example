package handler

import (
	"chatgpt-api-example/constant"
	"chatgpt-api-example/dto"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"net/http"
	"strings"
)

// TODO: implement the config loader for better security
const APIKey = "<YOUR API KEY>"
const Host = "https://api.openai.com"
const ChatGPTModel = "gpt-3.5-turbo"

type GrammarChecker interface {
	Check(w http.ResponseWriter, r *http.Request)
}

func NewGrammarCheckerHandler() GrammarChecker {
	client := resty.New().
		SetAuthToken(APIKey).
		SetHeader("Content-Type", "application/json").
		SetBaseURL(Host)

	return &GrammarCheckerImpl{
		client: client,
	}
}

type GrammarCheckerImpl struct {
	client *resty.Client
	// other dependencies
}

func (h *GrammarCheckerImpl) Check(w http.ResponseWriter, r *http.Request) {
	grammarCheckRequestDto := dto.GrammarCheckRequestDto{}

	if err := json.NewDecoder(r.Body).Decode(&grammarCheckRequestDto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	body := &dto.ChatGPTRequest{
		Model: ChatGPTModel,
		Messages: []dto.Message{
			{
				Role:    "system",
				Content: constant.GrammarCheckPromptSystemMessage,
			},
			{
				Role:    "user",
				Content: getGrammarCheckUserMessage(&grammarCheckRequestDto),
			},
		},
	}

	res, err := h.client.R().
		SetBody(body).
		Post("v1/chat/completions")

	if err != nil {
		log.Fatal("error occur while calling API:", err)
	}

	if res.StatusCode() >= 400 {
		log.Fatal("got error status code:", res.StatusCode())
	}

	// unmarshal to raw content
	var data map[string]interface{}
	if err := json.Unmarshal(res.Body(), &data); err != nil {
		log.Fatal("Error while decoding JSON response", err)
	}

	rawContent := data["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
	result, err := formatGrammarCheck(rawContent)

	if err != nil {
		log.Fatal("Error while format JSON to Struct", err)
	}

	if err := json.NewEncoder(w).Encode(result); err != nil {
		log.Fatal("Error while encoding response", err)
	}
}

const CodeBlockJSONStart = "```json"
const CodeBlockJSONEnd = "```"
const CodeBlockStartLen = len(CodeBlockJSONStart)
const CodeBlockEndLen = len(CodeBlockJSONEnd)

// formatGrammarCheck convert string to struct
// remove code block (sometime response include the md code block, it messes up the json unmarshal)
func formatGrammarCheck(content string) (*dto.GrammarCheckResponseDto, error) {
	startIndex := CodeBlockStartLen
	endIndex := len(content) - CodeBlockEndLen

	if len(content) < CodeBlockStartLen+CodeBlockEndLen {
		return nil, errors.New("content is too short to be a valid code block")
	}

	if content[:startIndex] == CodeBlockJSONStart && content[endIndex:] == CodeBlockJSONEnd {
		content = content[startIndex:endIndex]
	}

	content = strings.ReplaceAll(content, "\\\\\\", "\\")

	var result *dto.GrammarCheckResponseDto
	if err := json.Unmarshal([]byte(content), &result); err != nil {
		return nil, err
	}

	return result, nil
}

func getGrammarCheckUserMessage(grammarCheckDto *dto.GrammarCheckRequestDto) string {
	return fmt.Sprintf("Correct this paragraph: %s", grammarCheckDto.Paragraph)
}
