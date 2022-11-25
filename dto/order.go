package dto

type OrderUpdateDTO struct {
	ID     uint64 `json:"id" form:"id" binding:"required"`
	Status int    `json:"status" form:"status" binding:"required"`
	UserID uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
}

type OrderCreateDTO struct {
	UserID  uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
	Items   string `json:"items"  form:"items" binding:"required"`
	Total   int    `json:"total"  form:"total" binding:"required"`
	Status  int    `json:"status" default:"0"`
	Comment string `json:"comment" form:"comment"`
}
