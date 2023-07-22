package api

type WriterType string

const (
	ErrorW WriterType = "error"
	InfoW  WriterType = "info"
	DebugW WriterType = "debug"
)

type Response struct {
	StatusCode int    `json:"status_Code"`
	Message    string `json:"message"`
}
type WriterRequest struct {
	Writer WriterType `json:"writer"`
}
