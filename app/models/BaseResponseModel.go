package models

type BaseResponseModel struct {
	Status      int         `json:"status"`
	Message     string      `json:"message"`
	Total       uint        `json:"total,omitempty"`
	TotalPage   uint        `json:"total_page,omitempty"`
	CurrentPage uint        `json:"current_page,omitempty"`
	Data        interface{} `json:"data"`
	Success     bool        `json:"success"`
}
