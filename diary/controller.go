package diary

import (
	"log"

	"github.com/gin-gonic/gin"

	"my_diary/dto"
	"my_diary/model"
	"my_diary/utils"
	"my_diary/constant"
)

func RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/create", createDiary)

	r.POST("/get-num", getDiaryNum)
	r.POST("/list", getDiaryList)
	r.POST("/get", getDiary)

	r.PUT("/", updateDiary)

	r.DELETE("/", deleteDiary)
}

func createDiary(c *gin.Context) {
	var createDiaryReq CreateDiaryRequest
	err := c.ShouldBindJSON(&createDiaryReq)
	defer utils.LogError(&err)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
	}
	
	userId, exists := c.Get("userId")
	if !exists {
		dto.ErrorResponse(c, dto.WithMessage("未登录"), dto.WithCode(constant.NeedLogin))
		return
	}
	createDiaryReq.UserId = userId.(int)
	err = model.ExamineData(&createDiaryReq)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
		return
	}
	err = CreateDiaryService(createDiaryReq)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
		return
	}
	dto.SuccessResponse(c)
}

// 通过id获取
func getDiary(c *gin.Context) {
	var getDiaryReq GetDiaryRequest
	err := c.ShouldBindJSON(&getDiaryReq)
	defer utils.LogError(&err)

	userId, _ := c.Get("userId")
	getDiaryReq.UserId = userId.(int)
	err = model.ExamineData(&getDiaryReq)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
		return
	}
	resp, err := GetDiaryByDiaryIdService(getDiaryReq)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
		return
	}
	dto.SuccessResponse(c, dto.WithData(resp))
}

// 获取日记数量
func getDiaryNum(c *gin.Context) {
	var getDiaryNumReq GetDiaryNumRequest
	c.ShouldBindJSON(&getDiaryNumReq)
	
	userId, _ := c.Get("userId")
	getDiaryNumReq.UserId = userId.(int)

	resp, err := GetDiaryNumService(getDiaryNumReq)
	defer utils.LogError(&err)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
		return
	}
	dto.SuccessResponse(c, dto.WithData(resp))
}

func getDiaryList(c *gin.Context) {
	var getDiaryListReq GetDiaryListRequest
	err := c.ShouldBindJSON(&getDiaryListReq)
	defer utils.LogError(&err)

	userId, _ := c.Get("userId")
	getDiaryListReq.UserId = userId.(int)
	err = model.ExamineData(&getDiaryListReq)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
		return
	}
	resp, err := GetDiaryListService(getDiaryListReq)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
		return
	}
	dto.SuccessResponse(c, dto.WithData(resp))
}

func updateDiary(c *gin.Context) {
	var updateDiaryReq UpdateDiaryRequest
	err := c.ShouldBindJSON(&updateDiaryReq)
	defer utils.LogError(&err)

	userId, exists := c.Get("userId")
	if !exists {
		dto.ErrorResponse(c, dto.WithMessage("未登录"), dto.WithCode(constant.NeedLogin))
		return
	}
	updateDiaryReq.UserId = userId.(int)
	err = model.ExamineData(&updateDiaryReq)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
		return
	}
	err = UpdateDiaryService(updateDiaryReq)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
		return
	}
	dto.SuccessResponse(c)
}

func deleteDiary(c *gin.Context) {
	var deleteDiaryReq DeleteDiaryRequest
	err := c.ShouldBindJSON(&deleteDiaryReq)
	log.Println(deleteDiaryReq)
	defer utils.LogError(&err)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
		return
	}

	userId, exists := c.Get("userId")
	if !exists {
		dto.ErrorResponse(c, dto.WithMessage("未登录"), dto.WithCode(constant.NeedLogin))
		return
	}
	deleteDiaryReq.UserId = userId.(int)
	err = model.ExamineData(&deleteDiaryReq)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
		return
	}
	err = DeleteDiaryService(deleteDiaryReq)
	if err != nil {
		dto.ErrorResponse(c, dto.WithMessage(err.Error()))
		return
	}
	dto.SuccessResponse(c)
}

