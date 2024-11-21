package response

type ApiResponse struct {
	Message string      `json:"message"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type AutocomplateResponse struct {
	Value string `json:"value"`
	Label string `json:"label"`
}
