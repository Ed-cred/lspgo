package lsp

type DidOpenTextDocNotification struct {
	Notification
	Params DidOpenTextDocParams `json:"params"`
}

type DidOpenTextDocParams struct {
	TextDocument TextDocumentItem `json:"textDocument"`
}
