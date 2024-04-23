package request

type UserRequest struct {
	Name     string `json:"name" binding:"required,min=4,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,containsany=$!@#%&*"`
	Age      int32  `json:"age" binding:"required,numeric,min=12,max=100"`
}
