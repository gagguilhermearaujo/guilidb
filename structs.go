package main

type ErrorMessage struct {
	Error   string         `json:"error"`
	Example map[string]any `json:"example,omitempty"`
}

type DocumentRequest struct {
	Collection string         `json:"collection" xml:"collection" form:"collection"`
	Key        string         `json:"key" xml:"key" form:"key"`
	Data       map[string]any `json:"data" xml:"data" form:"data"`
}
type DocumentRequests struct {
	Documents []DocumentRequest `json:"documents" xml:"documents" form:"documents"`
}

type Responses struct {
	Documents []DocumentResponse `json:"documents" xml:"documents" form:"documents"`
}

type DocumentResponse struct {
	Collection string           `json:"collection" xml:"collection" form:"collection"`
	Key        string           `json:"key" xml:"key" form:"key"`
	StatusCode int              `json:"status_code" xml:"status_code" form:"status_code"`
	Message    string           `json:"message" xml:"message" form:"message"`
	Data       map[string]any   `json:"data,omitempty" xml:"data,omitempty" form:"data,omitempty"`
	Audit      []map[string]any `json:"audit,omitempty" xml:"audit,omitempty" form:"audit,omitempty"`
}
