package app

import (
	"github.com/semidark/kicli/internal/aiclient"
	"github.com/semidark/kicli/internal/configmanager"
	"github.com/semidark/kicli/internal/storage"
)

// Message types for Bubbletea communication as defined in docs/interfaces.md

// PTY-related messages
type PtyOutputMsg struct {
	Data []byte
}

type PtyExitedMsg struct {
	Err error
}

// AI-related messages
type AIResponseChunkMsg struct {
	Chunk string
}

type AIResponseCompleteMsg struct {
	Full string
	Role string
	Err  error
}

type CommandSuggestionChunkMsg struct {
	Chunk string
}

type CommandGeneratedMsg struct {
	Command string
	Error   error
}

type AIRequestSentMsg struct {
	RequestID string
}

// Application control messages
type ExecuteCommandInPtyMsg struct {
	Command string
}

type ErrorOccurredMsg struct {
	Err error
}

// Configuration and storage messages
type ConfigLoadedMsg struct {
	Cfg configmanager.AppConfig
	Err error
}

type HistoryLoadedMsg struct {
	Chat  []aiclient.ChatMessage
	Shell []storage.ShellCommandEntry
	Err   error
}
