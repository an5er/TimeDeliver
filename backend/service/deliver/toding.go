package deliver

import (
	"timedeliver/global"
	"timedeliver/model"
	"timedeliver/model/request"
	"timedeliver/model/response"
)

type TodoingService struct{}

var TodoingServiceApp = new(TodoingService)

func (todoingService *TodoingService) CreateTodoing(toding model.Todoing) (err error) {
	return global.DB.Create(toding).Error
}

func (todoingService *TodoingService) GetTodoingByMonth(time request.TodoingMonth, uid int) (err []response.Todoing) {
	global.Zap.Info("为用户:" + string(uid) + "查询" + time.DeadlineYear + "-" + time.DeadlineMonth)
	var todoArr []response.Todoing
	global.DB.Table("todoings").Select("is_done,deadline,thing,remark").
		Where("uid = ?", uid).Where("YEAR(deadline)=? AND MONTH(deadline)=?", time.DeadlineYear, time.DeadlineMonth).
		Find(&todoArr)

	return todoArr
}
