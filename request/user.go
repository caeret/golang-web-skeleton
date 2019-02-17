package request

type CreateUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
