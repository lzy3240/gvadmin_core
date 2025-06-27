package baseapi

import (
	vd "github.com/bytedance/go-tagexpr/v2/validator"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gvadmin_v3/core/baseapi/constructor"
	"gvadmin_v3/core/baseservice"
	"gvadmin_v3/core/util"
)

type Api struct {
	c *gin.Context
	r *CommonResp
}

// MountCtx 挂载上下文
func (a *Api) MountCtx(c *gin.Context) *Api {
	if a.c == nil {
		a.c = c
	}
	return a
}

func (a *Api) SetService(s *baseservice.Service) *Api {
	userId, _ := a.c.Get("userId")
	rid, _ := a.c.Get("requestId")
	dataScope, _ := a.c.Get("dataScope")
	deptId, _ := a.c.Get("deptId")
	//roleId, _ := a.c.Get("roleId")

	s.RequestID = rid.(string)
	s.DP = baseservice.DataPermission{
		DataScope: dataScope.(string),
		UserId:    util.AnyToInt(userId),
		DeptId:    util.AnyToInt(deptId),
		RoleId:    0,
	}
	return a
}

// GetUserFromCtx 获取当前操作用户id
func (a *Api) GetUserFromCtx() int {
	userId, b := a.c.Get("userId")
	if b {
		return util.AnyToInt(userId)
	} else {
		return 0
	}
}

// Bind 参数校验
func (a *Api) Bind(d interface{}, bindings ...binding.Binding) error {
	var err error
	if len(bindings) == 0 {
		bindings = constructor.Constructor.GetBindingForGin(d)
	}
	for i := range bindings {
		if bindings[i] == nil {
			err = a.c.ShouldBindUri(d)
		} else {
			err = a.c.ShouldBindWith(d, bindings[i])
		}
		if err != nil && err.Error() == "EOF" {
			err = nil
			continue
		}
		if err != nil {
			return err
		}
	}

	if err1 := vd.Validate(d); err1 != nil {
		return err1
	}

	return nil
}
