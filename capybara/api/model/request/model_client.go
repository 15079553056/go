package request

type ClientRegister struct{
	Email string `json:"email"`
	Phone string `json:"phone"`
	Password string `json:"password"`
}