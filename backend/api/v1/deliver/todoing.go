package deliver

import (
	"github.com/gin-gonic/gin"
	"time"
	"timedeliver/global"
	"timedeliver/model"
	"timedeliver/model/request"
	"timedeliver/model/response"
	"timedeliver/util"
)

type TodingApi struct{}

func (e *TodingApi) CreateToding(c *gin.Context) {
	var todoing request.Todoing

	err := c.ShouldBindJSON(&todoing)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	token := c.Request.Header.Get("x-token")
	j := util.NewJWT()
	claims, err := j.ParseToken(token)
	global.Zap.Info(claims.Username + "创建代办")

	var mtoding model.Todoing

	t, err := time.Parse("2006-01-02", todoing.Deadline)
	mtoding.Uid = claims.UID
	mtoding.Deadline = t
	mtoding.Is_Done = todoing.Is_Done
	mtoding.Remark = todoing.Remark
	mtoding.Thing = todoing.Thing

	err = todingService.CreateTodoing(mtoding)
	if err != nil {
		response.FailWithMessage("创建代办失败", c)
		global.Zap.Error("创建代办事项失败" + err.Error())
		return
	}
	response.OkWithMessage("创建代办成功", c)
}

func (e *TodingApi) GetTodingByMonth(c *gin.Context) {
	var time request.TodoingMonth

	err := c.ShouldBindJSON(&time)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	token := c.Request.Header.Get("x-token")
	j := util.NewJWT()
	claims, err := j.ParseToken(token)

	todoArr := todingService.GetTodoingByMonth(time, claims.UID)
	response.OkWithData(todoArr, c)
	return
}
