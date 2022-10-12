package request

type InputLetter struct {
	Description string `json:"description" binding:"required"`
	Link        string `json:"link" binding:"required"`
	Month       string `json:"month" binding:"required"`
}
