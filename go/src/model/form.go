package model

type AskForChangePasswordForm struct {
	Email string `json:"email" binding:"required,email"`
}

type AskForRegisterForm struct {
	Email string `json:"email" binding:"required,email"`
}

type RegisterForm struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,password"`
	Vrc      string `json:"vrc" binding:"required,vrc"`
}

type ChangePasswordForm struct {
	Email       string `json:"email" binding:"required,email"`
	NewPassword string `json:"new_password" binding:"required,password"`
	Vrc         string `json:"vrc" binding:"required,vrc"`
}

type LoginForm struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,password"`
}

type UpdateAvatarForm struct {
	AvatarBase64Data string `json:"avatar_base64_data" binding:"required,base64,avatar"`
}
