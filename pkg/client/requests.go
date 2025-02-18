package v1

type UpdateStockPrice struct {
	ID    uint    `json:"id" binding:"required"`
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}

//-----------------------------------------------------------
//------Users

type CreateUserRequest struct {
	Name  string `binding:"required" json:"name"`
	Email string `binding:"required,email" json:"email"`
}

type LoginUserRequest struct {
	Email string `binding:"required,email" json:"email"`
}
