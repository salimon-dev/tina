package openai

import "encoding/json"

type CompletionMessage struct {
	Role       string               `json:"role"`
	ToolCallId string               `json:"tool_call_id"`
	Content    string               `json:"content"`
	ToolCalls  []CompletionToolCall `json:"tool_calls"`
}

type CompletionToolCall struct {
	Id       string                     `json:"id"`
	Type     string                     `json:"type"`
	Function CompletionToolCallFunction `json:"function"`
}

type CompletionToolCallFunction struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

type CompletionFunction struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Strict      bool        `json:"strict"`
	Parameters  interface{} `json:"parameters"`
}

type CompletionTool struct {
	Type     string             `json:"type"`
	Function CompletionFunction `json:"function"`
}

type CompletionParams struct {
	Messages            []CompletionMessage `json:"messages"`
	Model               string              `json:"model"`
	MaxCompletionTokens int                 `json:"max_completion_tokens"`
	Temperature         float64             `json:"temperature"`
	Tools               []CompletionTool    `json:"tools"`
}

type EmbedParams struct {
	Input string `json:"input"`
	Model string `json:"model"`
}

type CompletionResponseChoice struct {
	Index   int32             `json:"index"`
	Message CompletionMessage `json:"message"`
}

type CompletionUsage struct {
	PromptTokens     int32 `json:"prompt_tokens"`
	CompletionTokens int32 `json:"completion_tokens"`
	TotalTokens      int32 `json:"total_tokens"`
}

type EmbeddingUsage struct {
	PromptTokens int32 `json:"prompt_tokens"`
	TotalTokens  int32 `json:"total_tokens"`
}

type CompletionResponse struct {
	Id      string                     `json:"id"`
	Object  string                     `json:"object"`
	Created int32                      `json:"created"`
	Model   string                     `json:"model"`
	Choices []CompletionResponseChoice `json:"choices"`
	Usage   CompletionUsage            `json:"usage"`
}

type EmbeddingData struct {
	Type      string    `json:"type"`
	Index     int       `json:"index"`
	Embedding []float64 `json:"embedding"`
}

type EmbeddingResponse struct {
	Object string          `json:"object"`
	Model  string          `json:"model"`
	Usage  EmbeddingUsage  `json:"usage"`
	Data   []EmbeddingData `json:"data"`
}

type Action struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Meta        string          `json:"meta"`
	Parameters  json.RawMessage `json:"parameters"`
	Vectors     []float64       `json:"vectors"`
}

type CompletionParsedResponse struct {
	Body string
}
