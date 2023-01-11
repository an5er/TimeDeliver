package response

import "timedeliver/model/request"

type UserResponse struct {
	User request.User `json:"user"`
}

type LoginResponse struct {
	User      request.User `json:"user"`
	Token     string       `json:"token"`
	ExpiresAt int64        `json:"expiresAt"`
}
