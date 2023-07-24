package model

import "time"

type RegisterReq struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResp struct {
	Token          string     `json:"token"`
	TokenExpiredAt *time.Time `json:"tokenExpiredAt"`
}

type AddWorkReq struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ListWorksReq struct {
	UserId   uint64 `json:"user_id"`
	Page     uint64 `json:"page"`
	PageSize uint64 `json:"page_size"`
}

type ListWorksResp struct {
	Works []Work `json:"works"`
	Count int64  `json:"count"`
}

type GetWorkResp struct {
	Work Work `json:"work"`
}

type GetProfileResp struct {
	User User `json:"user"`
}

type EditMyProfileReq struct {
	User User `json:"user"`
}
