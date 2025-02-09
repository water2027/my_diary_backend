package user

import (
	"errors"
)

type userModel struct {
	userId    int // 自增主键
	userEmail string
	username  string
	password  string
}

type LoginRequest struct {
	UserEmail string `json:"email"`
	Password  string `json:"password"`
}

func (lr *LoginRequest) Examine() error {
	if lr.UserEmail == "" || lr.Password == "" {
		return errors.New("email or password is empty")
	}
	return nil
}

type LoginResponse struct {
	UserEmail string `json:"email"`
	Username  string `json:"username"`
	Token     string `json:"token"`
}

type RegisterRequest struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
	Code      string `json:"code"`
}

func (rr *RegisterRequest) Examine() error {
	if rr.Email == "" || rr.Username == "" || rr.Password == "" || rr.Password2 == "" || rr.Code == "" {
		return errors.New("invalid attribute")
	}
	if rr.Password != rr.Password2 {
		return errors.New("passwords are not the same")
	}
	return nil
}

type RegisterResponse struct {
	LoginResponse
}

type UpdatePasswordRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
	Code      string `json:"code"`
}

func (upr *UpdatePasswordRequest) Examine() error {
	if upr.Email == "" || upr.Password == "" || upr.Password2 == "" || upr.Code == "" {
		return errors.New("invalid attribute")
	}
	if upr.Password != upr.Password2 {
		return errors.New("passwords are not the same")
	}
	return nil
}

//

type SendCodeRequest struct {
	Email string `json:"email"`
}

func (scr *SendCodeRequest) Examine() error {
	if scr.Email == "" {
		return errors.New("invalid email")
	}
	return nil
}

type SendCodeResponse string

type DeleteRequest struct {
	UserId int    `json:"-"`
	Code   string `json:"code"`
}

func (dr *DeleteRequest) Examine() error {
	if dr.UserId <= 0 || dr.Code == "" {
		return errors.New("invalid attribute")
	}
	return nil
}

//

type GetUserInfoRequest struct {
	UserId int `json:"-"`
}

func (gur *GetUserInfoRequest) Examine() error {
	if gur.UserId <= 0 {
		return errors.New("invalid token")
	}
	return nil
}

type GetUserInfoResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}
