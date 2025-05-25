package aiclient

// ChatMessage represents a single chat message
type ChatMessage struct {
	Role    string // "user", "assistant", "system"
	Content string
}

// ChatMessageChunk represents a streaming chunk of a chat response
type ChatMessageChunk struct {
	Content string
	IsFinal bool
	Error   error
}

// CommandSuggestionChunk represents a streaming chunk of a command suggestion
type CommandSuggestionChunk struct {
	Command string
	IsFinal bool
	Error   error
}
