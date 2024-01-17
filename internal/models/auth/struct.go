package auth

type UserModel struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password_generator" binding:"required"`
}
