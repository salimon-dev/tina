package openai

type CompletionMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type CompletionParams struct {
	Messages            []CompletionMessage `json:"messages"`
	Model               string              `json:"model"`
	MaxCompletionTokens int                 `json:"max_completion_tokens"`
	Temperature         float64             `json:"temperature"`
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

type CompletionResponse struct {
	Id      string                     `json:"id"`
	Object  string                     `json:"object"`
	Created int32                      `json:"created"`
	Model   string                     `json:"model"`
	Choices []CompletionResponseChoice `json:"choices"`
	Usage   CompletionUsage            `json:"usage"`
}
