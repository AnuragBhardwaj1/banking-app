package handler

type ListUserDao struct {
    Users []UserInfo `json:"users,omitempty"`
}

type UserInfo struct {
    Name    string  `json:"name,omitempty"`
    Email   string  `json:"email,omitempty"`
    ApiKey string  `json:"apiKey,omitempty"`
}
