package user

import (
	"github.com/gin-gonic/gin"

	"my_diary/dto"
	"my_diary/model"
	"my_diary/utils"
)

func RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/login", login)
	r.POST("/register", register)
	r.POST("/send-code", sendCode)

	r.PUT("/update-password", updatePassword)

	r.DELETE("/delete-user", deleteUser)
}

func login(c *gin.Context) {
	var loginReq LoginRequest
	err := c.ShouldBindJSON(&loginReq)
	defer utils.LogError(&err)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
		return
	}
	err = model.ExamineData(&loginReq)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
		return
	}
	loginResp, err := LoginService(loginReq)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
		return
	}
	dto.SuccessResponse(c, dto.WithData(loginResp))
}

func register(c *gin.Context) {
	var registerReq RegisterRequest
	err := c.ShouldBindJSON(&registerReq)
	defer utils.LogError(&err)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
		return
	}
	err = model.ExamineData(&registerReq)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
		return
	}
	registerResp, err := RegisterService(registerReq)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
		return
	}
	dto.SuccessResponse(c, dto.WithData(registerResp))
}

func updatePassword(c *gin.Context) {
	var updatePasswordReq UpdatePasswordRequest
	err := c.ShouldBindJSON(&updatePasswordReq)
	defer utils.LogError(&err)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
		return
	}
	err = UpdatePasswordService(updatePasswordReq)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
		return
	}
	dto.SuccessResponse(c, dto.WithMessage("重置成功！"))
}

func deleteUser(c *gin.Context) {
	deleteReq := DeleteRequest{}
	err := c.ShouldBindJSON(&deleteReq)
	defer utils.LogError(&err)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
		return
	}

	userId, exists := c.Get("userId")
	if !exists {
		dto.ErrorResponse(c, dto.WithMessage("invalid user"))
		return
	}
	deleteReq.UserId = userId.(int)

	err = model.ExamineData(&deleteReq)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
		return
	}
	err = DeleteService(deleteReq)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
	}
	dto.SuccessResponse(c, dto.WithMessage("感谢您的使用，希望我们还能再见！"))
}

func sendCode(c *gin.Context) {
	var sendCodeReq SendCodeRequest
	err := c.ShouldBindJSON(&sendCodeReq)
	defer utils.LogError(&err)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
		return
	}
	err = model.ExamineData(&sendCodeReq)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
		return
	}
	err = SendCodeService(sendCodeReq)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
		return
	}
	dto.SuccessResponse(c, dto.WithMessage("验证码已发送，请查收"))
}
