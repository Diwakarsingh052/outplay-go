package appErr



type ApplicationError struct {
	Msg        string `json:"message"`
	StatusCode int    `json:"status_code"`
	Status     string `json:"status"`
}
