package api

import (
	"TaskList/models/apireq"
	"TaskList/pkg/er"
	"TaskList/pkg/query_condition"
	"github.com/gin-gonic/gin"
)

// ListTask
// @Summary List Task 獲取全部 Task
// @Produce json
// @Accept json
// @Tags Task
// @Security Bearer
// @Param Bearer header string true "Admin JWT Token"
// @Param page query int true "Which Page"
// @Param per_page query int true "How many data per page"
// @Success 200 {object} apires.ListAllTask
// @Failure 400 {object} er.AppErrorMsg "{"code":"400400","message":"Wrong parameter format or invalid"}"
// @Failure 401 {object} er.AppErrorMsg "{"code":"400401","message":"Unauthorized"}"
// @Failure 500 {object} er.AppErrorMsg "{"code":"500000","message":"Database unknown error"}"
// @Router /v1/tasks [GET]
func ListTask(c *gin.Context) {
	req := &apireq.ListTask{}
	if err := c.BindQuery(req); err != nil {
		err = er.NewAppErr(400, er.ErrorParamInvalid, err.Error(), err)
		_ = c.Error(err)
		return
	}

	// need to change to what you want to query
	var reqCondition apireq.ListTaskQueryCondition
	if err := c.BindQuery(&reqCondition); err != nil {
		err = er.NewAppErr(400, er.ErrorParamInvalid, err.Error(), err)
		_ = c.Error(err)
		return
	}
	queryCondition := query_condition.QueryCondition{Condition: reqCondition}

	// need to dependency injection
	srv := BuildTaskSrv()
	res, err := srv.FindAll(req, queryCondition)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(200, res)
}

// GetTaskDetail
// @Summary GetTaskDetail 獲取Task詳細資訊
// @Produce json
// @Accept json
// @Tags Task
// @Security Bearer
// @Param Bearer header string true "Admin JWT Token"
// @Param id path int true "task_id"
// @Success 200 {object} apires.Task
// @Failure 400 {object} er.AppErrorMsg "{"code":"400400","message":"Wrong parameter format or invalid"}"
// @Failure 401 {object} er.AppErrorMsg "{"code":"400401","message":"Unauthorized"}"
// @Failure 404 {object} er.AppErrorMsg "{"code":"400401","message":"Data not found"}"
// @Failure 500 {object} er.AppErrorMsg "{"code":"500000","message":"Database unknown error"}"
// @Router /v1/tasks/{id} [GET]
func GetTaskDetail(c *gin.Context) {
	req := &apireq.GetTaskDetail{}
	err := c.BindUri(&req)
	if err != nil {
		err = er.NewAppErr(400, er.ErrorParamInvalid, err.Error(), err)
		_ = c.Error(err)
		return
	}
	// need to dependency injection
	srv := BuildTaskSrv()
	res, err := srv.FindOne(req)
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
// @Security Bearer
// @Param Bearer header string true "Admin JWT Token"
// @Param Body body apireq.CreateTask true "Request 新增 Task"
// @Success 200 {string} string "{}"
// @Failure 400 {object} er.AppErrorMsg "{"code":"400400","message":"Wrong parameter format or invalid"}"
// @Failure 401 {object} er.AppErrorMsg "{"code":"400401","message":"Unauthorized"}"
// @Failure 500 {object} er.AppErrorMsg "{"code":"500000","message":"Database unknown error"}"
// @Router /v1/tasks [POST]
func CreateTask(c *gin.Context) {
	req := &apireq.CreateTask{}
	if err := c.Bind(req); err != nil {
		err = er.NewAppErr(400, er.ErrorParamInvalid, err.Error(), err)
		_ = c.Error(err)
		return
	}
	// need to dependency injection
	srv := BuildTaskSrv()
	err := srv.Create(req)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(200, struct{}{})
}

// UpdateTask
// @Summary Update Task 修改Task
// @Produce json
// @Accept json
// @Tags Task
// @Security Bearer
// @Param Bearer header string true "Admin JWT Token"
// @Param id path int true "task_id"
// @Param Body body apireq.UpdateTask true "Request 修改 Task"
// @Success 200 {string} string "{}"
// @Failure 400 {object} er.AppErrorMsg "{"code":"400400","message":"Wrong parameter format or invalid"}"
// @Failure 401 {object} er.AppErrorMsg "{"code":"400401","message":"Unauthorized"}"
// @Failure 404 {object} er.AppErrorMsg "{"code":"400401","message":"Data not found"}"
// @Failure 500 {object} er.AppErrorMsg "{"code":"500000","message":"Database unknown error"}"
// @Router /v1/tasks/{id} [PUT]
func UpdateTask(c *gin.Context) {
	req := &apireq.UpdateTask{}

	_ = c.Bind(req)
	err := c.BindUri(&req)
	if err != nil {
		err = er.NewAppErr(400, er.ErrorParamInvalid, err.Error(), err)
		_ = c.Error(err)
		return
	}
	// need to dependency injection
	srv := BuildTaskSrv()
	err = srv.Update(req)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(200, struct{}{})
}

// DeleteTask
// @Summary Delete Task 刪除Task
// @Produce json
// @Accept json
// @Tags Task
// @Security Bearer
// @Param Bearer header string true "Admin JWT Token"
// @Param id path int true "task_id"
// @Success 200 {string} string "{}"
// @Failure 400 {object} er.AppErrorMsg "{"code":"400400","message":"Wrong parameter format or invalid"}"
// @Failure 401 {object} er.AppErrorMsg "{"code":"400401","message":"Unauthorized"}"
// @Failure 404 {object} er.AppErrorMsg "{"code":"400401","message":"Data not found"}"
// @Failure 500 {object} er.AppErrorMsg "{"code":"500000","message":"Database unknown error"}"
// @Router /v1/tasks/{id} [DELETE]
func DeleteTask(c *gin.Context) {
	req := &apireq.DeleteTask{}
	err := c.BindUri(&req)
	if err != nil {
		err = er.NewAppErr(400, er.ErrorParamInvalid, err.Error(), err)
		_ = c.Error(err)
		return
	}
	// need to dependency injection
	srv := BuildTaskSrv()
	err = srv.Delete(req)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(200, struct{}{})
}
