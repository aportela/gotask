package middlewares

type errorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
