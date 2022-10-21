package dto

type ProductUpdateDTO struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	Name        string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Price       int    `json:"price" form:"price" binding:"required"`
	Status      bool   `json:"status"  form:"status"`
}

type ProductCreateDTO struct {
	Name        string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Price       int    `json:"price" form:"price" binding:"required"`
	Status      bool   `json:"status"  form:"status"`
}
