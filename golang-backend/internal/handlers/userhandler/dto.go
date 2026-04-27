package userhandler

type addUserRequest struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsSuperUser bool   `json:"isSuperUser"`
}

type updateUserRequest struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	Password    *string `json:"password,omitempty"`
	IsSuperUser bool    `json:"isSuperUser"`
}

type userResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   *int64 `json:"updatedAt,omitempty"`
	IsSuperUser bool   `json:"isSuperUser"`
	AvatarURL   string `json:"avatar"`
}

type addUserResponse struct {
	User userResponse `json:"user"`
}

type updateUserResponse struct {
	User userResponse `json:"user"`
}

type getUserResponse struct {
	User userResponse `json:"user"`
}

type searchUsersResponse struct {
	Users []userResponse `json:"users"`
}
