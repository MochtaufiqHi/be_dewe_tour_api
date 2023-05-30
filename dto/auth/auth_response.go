package authdto

type AuthResponse struct {
	Fullname string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Fullname string `json:"name"`
	Email    string `json:"email"`
	// Password string `json:"password"`
	Token string `json:"token"`
}

type RegisterResponse struct {
	// Fullname string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	// Token string `json:"token"`
}
