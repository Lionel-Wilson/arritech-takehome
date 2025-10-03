package dto

type GetUserResponse struct {
	User User `json:"user"`
}

type GetUsersResponse struct {
	Users []User `json:"users"`
	Page  int    `json:"page"`
	Size  int    `json:"size"`
	Total int64  `json:"total"`
}

type User struct {
	ID          uint   `json:"id"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Age         int    `json:"age"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phonenumber"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
