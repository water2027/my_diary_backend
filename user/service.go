package user

import (
	"errors"
	"time"

	"context"

	"my_diary/utils"
)

func LoginService(loginReq LoginRequest) (LoginResponse, error) {
	user := userModel{
		userEmail: loginReq.UserEmail,
	}
	var loginResp LoginResponse
	var store userHandler = &user
	hasRegister := store.findUser()
	if !hasRegister {
		return loginResp, errors.New("wrong email or password")
	}
	if !utils.CheckPassword(user.password, loginReq.Password) {
		return loginResp, errors.New("invalid password")
	}
	token, err := utils.GenerateToken(user.userId)
	if err != nil {
		return loginResp, errors.New("token error")
	}
	loginResp.UserEmail = user.userEmail
	loginResp.Username = user.username
	loginResp.Token = token
	return loginResp, nil
}

func RegisterService(ctx context.Context, registerReq RegisterRequest) (RegisterResponse, error) {
	var registerResp RegisterResponse
	var err error
	// TODO: 数据库查询Code是否正确
	code, err := getCode(ctx, registerReq.Email)
	if err != nil {
		return registerResp, errors.New("验证码过期")
	}
	if code != registerReq.Code {
		return registerResp, errors.New("验证码错误")
	}
	user := userModel{
		userEmail: registerReq.Email,
		username:  registerReq.Username,
	}
	var store userHandler = &user
	// 查询邮箱是否已经注册
	hasRegister := store.findUser()
	if hasRegister {
		return registerResp, errors.New("email has been registered")
	}

	// 插入用户数据
	user.password, err = utils.HashPassword(registerReq.Password)
	if err != nil {
		return registerResp, err
	}
	err = store.insert()
	if err != nil {
		return registerResp, err
	}

	// 获取用户ID
	ok := store.findUser()
	if !ok {
		return registerResp, errors.New("数据库错误")
	}

	// 生成token
	token, err := utils.GenerateToken(user.userId)
	if err != nil {
		return registerResp, err
	}
	registerResp.UserEmail = user.userEmail
	registerResp.Username = user.username
	registerResp.Token = token
	deleteCode(ctx, user.userEmail)
	return registerResp, nil
}

func UpdatePasswordService(ctx context.Context, updatePasswordReq UpdatePasswordRequest) error {
	// TODO: 数据库查询Code是否正确
	code, err := getCode(ctx, updatePasswordReq.Email)
	if err != nil {
		return errors.New("验证码过期")
	}
	if code != updatePasswordReq.Code {
		return errors.New("验证码错误")
	}
	user := userModel{
		userEmail: updatePasswordReq.Email,
	}
	var store userHandler = &user
	hasRegister := store.findUser()
	if !hasRegister {
		return errors.New("unregister")
	}
	hashPassword, err := utils.HashPassword(updatePasswordReq.Password)
	if err != nil {
		return err
	}
	user.password = hashPassword
	err = store.updateUser()
	if err != nil {
		return err
	}
	deleteCode(ctx, user.userEmail)
	return nil
}

func DeleteService(ctx context.Context, deleteReq DeleteRequest) error {
	user := userModel{
		userId: deleteReq.UserId,
	}
	var store userHandler = &user
	hasFound := store.findUser()
	if !hasFound {
		return errors.New("未找到用户")
	}

	code, err := getCode(ctx, user.userEmail)
	if err != nil {
		return errors.New("请发送验证码")
	}
	if code != deleteReq.Code {
		return errors.New("验证码错误")
	}
	err = store.deleteUser()
	deleteCode(ctx, user.userEmail)
	return err
}

func SendCodeService(ctx context.Context, sendCodeReq SendCodeRequest) error {
	code := utils.GetRandomCode()
	body := "您的验证码是：" + code
	err := utils.SendMail(sendCodeReq.Email, "diary验证码", body)
	if err != nil {
		return err
	}
	// TODO: 将验证码存入数据库
	err = setCode(ctx, sendCodeReq.Email, code, time.Minute * 5)
	return err
}

func GetUserInfoService(getUserInfoReq GetUserInfoRequest) (GetUserInfoResponse, error) {
	user := userModel{
		userId: getUserInfoReq.UserId,
	}
	var store userHandler = &user
	success := store.findUser()
	if !success {
		return GetUserInfoResponse{}, errors.New("user not found")
	}
	return GetUserInfoResponse{
		Email:    user.userEmail,
		Username: user.username,
	}, nil
}
