package dto

type SignInDTO struct {
	Email    string `json:"email" binding:"required,email,max=64" example:"kill-77@mail.ru"`
	Password string `json:"password" binding:"required,min=6,max=64" example:"qwerty"`
}

type SignUpDTO struct {
	FirstName string `json:"firstName" binding:"required,min=2,max=64" example:"alex"`
	LastName  string `json:"lastName" binding:"required,min=2,max=64" example:"cross"`
	Email     string `json:"email" binding:"required,email,max=64" example:"kill-77@mail.ru"`
	Password  string `json:"password" binding:"required,min=6,max=64" example:"qwerty"`
}
