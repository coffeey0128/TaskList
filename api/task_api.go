package api

import (
	"TaskList/models/apireq"
	"TaskList/pkg/er"

	"github.com/gin-gonic/gin"
)

// ListTask
// @Summary List Task 獲取全部 Task
// @Produce json
// @Accept json
// @Tags Task
// @Success 200 {object} []apires.Task
// @Failure 400 {object} er.AppErrorMsg "{"code":"400400","message":"Wrong parameter format or invalid"}"
// @Failure 401 {object} er.AppErrorMsg "{"code":"400401","message":"Unauthorized"}"
// @Failure 500 {object} er.AppErrorMsg "{"code":"500000","message":"Database unknown error"}"
// @Router /tasks [GET]
func ListTask(c *gin.Context) {
	srv := BuildTaskSrv()
	res, err := srv.FindAll()
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(200, res)
}

// CreateTask
// @Summary Create Task 新增Task
// @Produce json
// @Accept json
// @Tags Task
// @Param Body body apireq.CreateTask true "Request 新增 Task"
// @Success 200 {string} models.Task
// @Failure 400 {object} er.AppErrorMsg "{"code":"400400","message":"Wrong parameter format or invalid"}"
// @Failure 401 {object} er.AppErrorMsg "{"code":"400401","message":"Unauthorized"}"
// @Failure 500 {object} er.AppErrorMsg "{"code":"500000","message":"Database unknown error"}"
// @Router /tasks [POST]
func CreateTask(c *gin.Context) {
	req := &apireq.CreateTask{}
	if err := c.Bind(req); err != nil {
		err = er.NewAppErr(400, er.ErrorParamInvalid, err.Error(), err)
		_ = c.Error(err)
		return
	}
	srv := BuildTaskSrv()
	res, err := srv.Create(req)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(201, res)
}

// UpdateTask
// @Summary Update Task 修改Task
// @Produce json
// @Accept json
// @Tags Task
// @Security Bearer
// @Param id path int true "task_id"
// @Param Body body apireq.UpdateTask true "Request 修改 Task"
// @Success 200 {string} models.Task
// @Failure 400 {object} er.AppErrorMsg "{"code":"400400","message":"Wrong parameter format or invalid"}"
// @Failure 401 {object} er.AppErrorMsg "{"code":"400401","message":"Unauthorized"}"
// @Failure 404 {object} er.AppErrorMsg "{"code":"400401","message":"Data not found"}"
// @Failure 500 {object} er.AppErrorMsg "{"code":"500000","message":"Database unknown error"}"
// @Router /tasks/{id} [PUT]
func UpdateTask(c *gin.Context) {
	req := &apireq.UpdateTask{}

	_ = c.Bind(req)
	err := c.BindUri(&req)
	if err != nil {
		err = er.NewAppErr(400, er.ErrorParamInvalid, err.Error(), err)
		_ = c.Error(err)
		return
	}
	srv := BuildTaskSrv()
	res, err := srv.Update(req)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(200, res)
}

// DeleteTask
// @Summary Delete Task 刪除Task
// @Produce json
// @Accept json
// @Tags Task
// @Param id path int true "task_id"
// @Success 200 {string} string "{}"
// @Failure 400 {object} er.AppErrorMsg "{"code":"400400","message":"Wrong parameter format or invalid"}"
// @Failure 401 {object} er.AppErrorMsg "{"code":"400401","message":"Unauthorized"}"
// @Failure 404 {object} er.AppErrorMsg "{"code":"400401","message":"Data not found"}"
// @Failure 500 {object} er.AppErrorMsg "{"code":"500000","message":"Database unknown error"}"
// @Router /tasks/{id} [DELETE]
func DeleteTask(c *gin.Context) {
	req := &apireq.DeleteTask{}
	err := c.BindUri(&req)
	if err != nil {
		err = er.NewAppErr(400, er.ErrorParamInvalid, err.Error(), err)
		_ = c.Error(err)
		return
	}
	srv := BuildTaskSrv()
	err = srv.Delete(req)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(200, struct{}{})
}
