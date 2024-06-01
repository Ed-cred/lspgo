package lsp

type TextDocDidChangeNotification struct {
	Notification
	Params DidChangeTextDocParams `json:"params"`
}

type DidChangeTextDocParams struct {
	TextDocument   VersionTextDocumentIdentifier `json:"textDocument"`
	ContentChanges []TextDocContentChangeEvent   `json:"contentChanges"`
}

type TextDocContentChangeEvent struct {
	Text string `json:"text"`
}
