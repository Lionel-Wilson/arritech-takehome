package dto

type GetUserResponse struct {
	User User `json:"user"`
}

type GetUsersResponse struct {
	Users []User `json:"users"`
}

type User struct {
	ID        uint   `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
