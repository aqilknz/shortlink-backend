package dto

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email" example:"aqil@c.com"`
	Password string `json:"password" binding:"required,min=8" example:"pass1234"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"aqil@c.com"`
	Password string `json:"password" binding:"required,min=8" example:"pass1234"`
}

type AuthResponse struct {
	Token string       `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZW1haWwiOiJ1c2VyQGV4YW1wbGUuY29tIiwiaWF0IjoxNjE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"`
	User  UserResponse `json:"user"`
}

type UserResponse struct {
	ID        int     `json:"id" example:"1"`
	Email     string  `json:"email" example:"aqil@c.com"`
	FullName  *string `json:"full_name" example:"John Doe"`
	AvatarURL *string `json:"avatar_url" example:"https://example.com/avatars/johndoe.png"`
}
