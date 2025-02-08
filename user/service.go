package user

import (
	"errors"

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

func RegisterService(registerReq RegisterRequest) (RegisterResponse, error) {
	var registerResp RegisterResponse
	var err error
	// TODO: 数据库查询Code是否正确
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
	return registerResp, nil
}

func UpdatePasswordService(updatePasswordReq UpdatePasswordRequest) error {
	// TODO: 数据库查询Code是否正确
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
	return nil
}

func DeleteService(deleteReq DeleteRequest) error {
	user := userModel{
		userId: deleteReq.UserId,
	}
	var store userHandler = &user
	err := store.deleteUser()
	
	return err
}

func SendCodeService(sendCodeReq SendCodeRequest) error {
	code := utils.GetRandomCode()
	body := "您的验证码是：" + code
	err := utils.SendMail(sendCodeReq.Email, "diary验证码", body)
	if err != nil {
		return err
	}
	// TODO: 将验证码存入数据库
	return nil
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