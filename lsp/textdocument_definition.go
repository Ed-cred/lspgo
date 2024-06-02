package lsp

type DefinitionRequest struct {
	Request
	Params DefinitonParams `json:"params"`
}

type DefinitonParams struct {
	TextDocumentPositionParams
}

type DefinitonResponse struct {
	Response
	Result Location `json:"result"`
}
