package models

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginRequestCode struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Code     int    `json:"code"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthInfo struct {
	UserID   string `json:"user_id"`
	UserRole string `json:"user_role"`
}

type RegisterRequest struct {
	Mail string `json:"mail"`
}

type RegisterConfirmRequest struct {
	AddTeacher AddTeacher
	Code       int `json:"code"`
}


type RegisterOTPRequest struct {
	Teacher Teacher
	Code       int `json:"code"`
}