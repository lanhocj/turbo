package lpc

type UserLPCServiceObject struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     int    `json:"role"`
}
